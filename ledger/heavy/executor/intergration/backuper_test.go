//
// Copyright 2019 Insolar Technologies GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// +build slowtest

package intergration

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/insolar/insolar/configuration"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/store"
	"github.com/insolar/insolar/ledger/heavy/executor"
	"github.com/insolar/insolar/network/storage"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

type testKey struct {
	id uint64
}

func (t *testKey) ID() []byte {
	bs := make([]byte, 8)
	binary.PutUvarint(bs, t.id)
	return bs
}

func (t *testKey) Scope() store.Scope {
	return store.ScopeJetDrop
}

func makeBackuperConfig(t *testing.T, prefix string) configuration.Backup {

	cfg := configuration.Backup{
		ConfirmFile:      "BACKUPED",
		MetaInfoFile:     "META.json",
		TargetDirectory:  "/tmp/BKP/TARGET/" + prefix,
		TmpDirectory:     "/tmp/BKP/TMP",
		DirNameTemplate:  "pulse-%d",
		BackupWaitPeriod: 60,
		BackupFile:       "incr.bkp",
		Enabled:          true,
	}

	err := os.MkdirAll(cfg.TargetDirectory, 0777)
	require.NoError(t, err)
	err = os.MkdirAll(cfg.TmpDirectory, 0777)
	require.NoError(t, err)

	return cfg
}

func clearData(t *testing.T, cfg configuration.Backup) {
	err := os.RemoveAll(cfg.TargetDirectory)
	require.NoError(t, err)
}

func TestBackuper(t *testing.T) {
	cfg := makeBackuperConfig(t, t.Name())
	defer clearData(t, cfg)

	tmpdir, err := ioutil.TempDir("", "bdb-test-")
	defer os.RemoveAll(tmpdir)
	require.NoError(t, err)

	db, err := store.NewBadgerDB(tmpdir)
	require.NoError(t, err)
	defer db.Stop(context.Background())

	bm, err := executor.NewBackupMaker(context.Background(), db, cfg, insolar.GenesisPulse.PulseNumber)
	require.NoError(t, err)

	savedKeys := make(map[store.Key]insolar.PulseNumber, 0)

	var stopWriting uint32
	sgWriteStopped := sync.WaitGroup{}
	sgWriteStopped.Add(1)

	testPulse := insolar.GenesisPulse.PulseNumber + insolar.PulseNumber(rand.Int()%20000+1)
	// writing data to db
	go func() {
		for i := 0; i < 2000000; i++ {
			if atomic.LoadUint32(&stopWriting) != 0 {
				break
			}
			key := &testKey{id: uint64(i)}
			value := testPulse + insolar.PulseNumber(i)
			err := db.Set(key, value.Bytes())
			require.NoError(t, err)
			savedKeys[key] = value
			require.NoError(t, err)
			time.Sleep(time.Duration(rand.Int()%10) * time.Millisecond)
		}
		sgWriteStopped.Done()
	}()

	wgBackup := sync.WaitGroup{}
	numIterations := 5

	wgBackup.Add(numIterations)
	// doing backups
	go func() {
		for i := 0; i < numIterations; i++ {
			err := bm.MakeBackup(context.Background(), testPulse+insolar.PulseNumber(i))
			require.NoError(t, err)
			wgBackup.Done()
			time.Sleep(time.Duration(rand.Int()%1000) * time.Millisecond)
		}
	}()

	// creating backup confirmation files
	go func() {
		for i := 0; i < numIterations+1; i++ {
			time.Sleep(2 * time.Second)

			backupConfirmFile := filepath.Join(makeCurrentBkpDir(cfg, testPulse+insolar.PulseNumber(i)), cfg.ConfirmFile)
			for true {
				fff, err := os.Create(backupConfirmFile)
				if err != nil && strings.Contains(err.Error(), "no such file or directory") {
					time.Sleep(time.Millisecond * 200)
					fmt.Printf("%s not created yet\n", backupConfirmFile)
					continue
				}
				require.NoError(t, err)
				require.NoError(t, fff.Close())
				break
			}
		}
	}()

	// wait for all backups done
	wgBackup.Wait()
	// stop writing to db
	atomic.StoreUint32(&stopWriting, 1)
	// wait for stopping
	sgWriteStopped.Wait()

	// final backup to collect all rest records
	err = bm.MakeBackup(context.Background(), testPulse+insolar.PulseNumber(numIterations))
	require.NoError(t, err)

	// check backup hashes
	checkBackupMetaInfo(t, cfg, numIterations, testPulse)

	// load all backups and check all records
	{
		recovTmpDir, err := ioutil.TempDir("", "bdb-test-")
		defer os.RemoveAll(recovTmpDir)
		require.NoError(t, err)
		recoveredDB, err := makeRawBadger(recovTmpDir)
		require.NoError(t, err)
		defer recoveredDB.Close()

		for i := 0; i < numIterations+1; i++ {
			bkpFileName := filepath.Join(
				cfg.TargetDirectory,
				fmt.Sprintf(cfg.DirNameTemplate, testPulse+insolar.PulseNumber(i)),
				cfg.BackupFile,
			)
			bkpFile, err := os.Open(bkpFileName)
			require.NoError(t, err)
			err = recoveredDB.Load(bkpFile, 2)
			require.NoError(t, err)
		}

		require.NotEqual(t, 0, len(savedKeys))

		for k, v := range savedKeys {
			gotRawValue, err := getFromDB(recoveredDB, k)
			require.NoError(t, err)
			gotPulseNumber := insolar.NewPulseNumber(gotRawValue)
			require.Equal(t, v, gotPulseNumber)
		}
	}

}

func makeRawBadger(dir string) (*badger.DB, error) {
	dir, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}

	ops := badger.DefaultOptions(dir)
	bdb, err := badger.Open(ops)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open badger")
	}

	return bdb, nil
}

func getFromDB(db *badger.DB, key store.Key) (value []byte, err error) {
	fullKey := append(key.Scope().Bytes(), key.ID()...)

	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(fullKey)
		if err != nil {
			return err
		}
		value, err = item.ValueCopy(nil)
		return err
	})

	if err != nil {
		if err == badger.ErrKeyNotFound {
			return nil, storage.ErrNotFound
		}
		return nil, err
	}

	return
}

func makeCurrentBkpDir(cfg configuration.Backup, pulse insolar.PulseNumber) string {
	return filepath.Join(cfg.TargetDirectory, fmt.Sprintf(cfg.DirNameTemplate, pulse))
}

func calculateFileHash(t *testing.T, fileName string) string {
	f, err := os.Open(fileName)
	require.NoError(t, err)
	defer f.Close()
	hasher := sha256.New()
	_, err = io.Copy(hasher, f)
	require.NoError(t, err)

	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func checkBackupMetaInfo(t *testing.T, cfg configuration.Backup, numIterations int, testPulse insolar.PulseNumber) {
	for i := 0; i < numIterations+1; i++ {
		currentPulse := testPulse + insolar.PulseNumber(i)
		currentBkpDir := makeCurrentBkpDir(cfg, currentPulse)
		metaInfo := filepath.Join(currentBkpDir, cfg.MetaInfoFile)
		raw, err := ioutil.ReadFile(metaInfo)
		require.NoError(t, err)

		bi := executor.BackupInfo{}
		err = json.Unmarshal(raw, &bi)
		require.NoError(t, err)

		// check file hash
		bkpFile := filepath.Join(currentBkpDir, cfg.BackupFile)
		md5sum := calculateFileHash(t, bkpFile)
		require.Equal(t, md5sum, bi.SHA256)

		// check pulse
		require.Equal(t, currentPulse, bi.Pulse)
	}
}
