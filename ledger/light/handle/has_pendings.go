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

package handle

import (
	"context"

	"github.com/insolar/insolar/insolar/flow"
	"github.com/insolar/insolar/insolar/payload"
	"github.com/insolar/insolar/ledger/light/proc"
	"github.com/pkg/errors"
)

type HasPendings struct {
	dep    *proc.Dependencies
	meta   payload.Meta
	passed bool
}

func NewHasPendings(dep *proc.Dependencies, meta payload.Meta, passed bool) *HasPendings {
	return &HasPendings{
		dep:    dep,
		meta:   meta,
		passed: passed,
	}
}

func (s *HasPendings) Present(ctx context.Context, f flow.Flow) error {
	msg := payload.HasPendings{}
	err := msg.Unmarshal(s.meta.Payload)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal HasPendings message")
	}

	passIfNotExecutor := !s.passed
	jet := proc.NewFetchJet(msg.ObjectID, flow.Pulse(ctx), s.meta, passIfNotExecutor)
	s.dep.FetchJet(jet)
	if err := f.Procedure(ctx, jet, true); err != nil {
		if err == proc.ErrNotExecutor && passIfNotExecutor {
			return nil
		}
		return err
	}

	objJetID := jet.Result.Jet

	hot := proc.NewWaitHot(objJetID, flow.Pulse(ctx), s.meta)
	s.dep.WaitHot(hot)
	if err := f.Procedure(ctx, hot, false); err != nil {
		return err
	}

	// To ensure, that we have the index. Because index can be on a heavy node.
	// If we don't have it and heavy does, UpdateObject fails because it should update light's index state
	getIndex := proc.NewEnsureIndex(msg.ObjectID, objJetID, s.meta)
	s.dep.EnsureIndex(getIndex)
	if err := f.Procedure(ctx, getIndex, false); err != nil {
		return err
	}

	hasPendings := proc.NewHasPendings(s.meta, msg.ObjectID)
	s.dep.HasPendings(hasPendings)
	if err := f.Procedure(ctx, hasPendings, false); err != nil {
		return err
	}

	return nil
}
