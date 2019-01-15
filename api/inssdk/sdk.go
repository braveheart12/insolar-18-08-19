/*
 *    Copyright 2019 INS Ecosystem
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package inssdk

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"sync"

	"github.com/insolar/insolar/api/requester"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/platformpolicy"
	"github.com/insolar/insolar/testutils"
	"github.com/pkg/errors"
)

type response struct {
	Error   string
	Result  interface{}
	TraceID string
}

type ringBuffer struct {
	sync.Mutex
	urls   []string
	cursor int
}

type memberKeys struct {
	Private string `json:"private_key"`
	Public  string `json:"public_key"`
}

func (rb *ringBuffer) Next() string {
	rb.Lock()
	defer rb.Unlock()
	rb.cursor++
	if rb.cursor >= len(rb.urls) {
		rb.cursor = 0
	}
	return rb.urls[rb.cursor]
}

type SDK struct {
	apiUrls    *ringBuffer
	rootMember *requester.UserConfigJSON
}

func NewSDK(urls []string, rootMemberKeysPath string) (*SDK, error) {
	buffer := &ringBuffer{urls: urls}

	rawConf, err := ioutil.ReadFile(rootMemberKeysPath)
	if err != nil {
		err = errors.Wrap(err, "[ NewSDK ] can't read keys from file")
		return nil, err
	}

	keys := memberKeys{}
	err = json.Unmarshal(rawConf, &keys)
	if err != nil {
		err = errors.Wrap(err, "[ NewSDK ] can't unmarshal keys")
		return nil, err
	}

	response, err := requester.Info(buffer.Next())
	if err != nil {
		err = errors.Wrap(err, "[ NewSDK ] can't get info")
		return nil, err
	}

	rootMember, err := requester.CreateUserConfig(response.RootMember, keys.Private)
	if err != nil {
		err = errors.Wrap(err, "[ NewSDK ] can't create user config")
		return nil, err
	}

	return &SDK{
		apiUrls:    buffer,
		rootMember: rootMember,
	}, nil

}

func (sdk *SDK) sendRequest(ctx context.Context, method string, params []interface{}, userCfg *requester.UserConfigJSON) ([]byte, error) {
	reqCfg := &requester.RequestConfigJSON{
		Params: params,
		Method: method,
	}

	body, err := requester.Send(ctx, sdk.apiUrls.Next(), userCfg, reqCfg)
	if err != nil {
		errors.Wrap(err, "[ sendRequest ] can not send request")
	}

	return body, nil
}

func (sdk *SDK) getResponse(body []byte) (*response, error) {
	res := &response{}
	err := json.Unmarshal(body, &res)
	if err != nil {
		return nil, errors.Wrap(err, "[ getResponse ] problems with unmarshal response")
	}

	return res, nil
}

func (sdk *SDK) CreateMember() (*Member, string, error) {
	ctx := inslogger.ContextWithTrace(context.Background(), "CreateMember")
	memberName := testutils.RandomString()
	ks := platformpolicy.NewKeyProcessor()

	privateKey, err := ks.GeneratePrivateKey()
	if err != nil {
		return nil, "", errors.Wrap(err, "[ CreateMember ] can't generate private key")
	}

	privateKeyStr, err := ks.ExportPrivateKeyPEM(privateKey)
	if err != nil {
		return nil, "", errors.Wrap(err, "[ CreateMember ] can't export private key")
	}

	memberPubKeyStr, err := ks.ExportPublicKeyPEM(ks.ExtractPublicKey(privateKeyStr))
	if err != nil {
		return nil, "", errors.Wrap(err, "[ CreateMember ] can't extract public key")
	}

	params := []interface{}{memberName, string(memberPubKeyStr)}
	body, err := sdk.sendRequest(ctx, "CreateMember", params, sdk.rootMember)
	if err != nil {
		return nil, "", errors.Wrap(err, "[ CreateMember ] can't send request")
	}

	response, err := sdk.getResponse(body)
	if err != nil {
		return nil, "", errors.Wrap(err, "[ CreateMember ] can't get response")
	}

	if response.Error != "" {
		return nil, response.TraceID, errors.New(response.Error)
	}

	return NewMember(response.Result.(string), string(privateKeyStr)), response.TraceID, nil
}

func (sdk *SDK) Transfer(ctx context.Context, amount uint, from *Member, to *Member) (string, string, error) {
	params := []interface{}{amount, to.Ref}
	config, err := requester.CreateUserConfig(from.Ref, to.PK)

	body, err := sdk.sendRequest(ctx, "Transfer", params, config)
	if err != nil {
		return "", "", err
	}
	response, err := sdk.getResponse(body)

	if response.Error != "" {
		return "", response.TraceID, errors.New(response.Error)
	}

	return response.Result.(string), response.TraceID, nil
}
