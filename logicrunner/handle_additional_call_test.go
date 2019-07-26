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

package logicrunner

import (
	"testing"
	"time"

	"github.com/gojuno/minimock"
	"github.com/stretchr/testify/require"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/bus"
	"github.com/insolar/insolar/insolar/flow"
	"github.com/insolar/insolar/insolar/gen"
	"github.com/insolar/insolar/insolar/message"
	"github.com/insolar/insolar/insolar/record"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/testutils"
)

func TestHandleAdditionalCallFromPreviousExecutor_Present(t *testing.T) {
	table := []struct {
		name                           string
		clarifyPendingStateResult      error
		startQueueProcessorResult      error
		expectedStartQueueProcessorCtr int32
		mocks                          func(t minimock.Tester) (*HandleAdditionalCallFromPreviousExecutor, flow.Flow)
		error                          bool
	}{
		{
			name: "Happy path",
			mocks: func(t minimock.Tester) (*HandleAdditionalCallFromPreviousExecutor, flow.Flow) {
				obj := gen.Reference()
				reqRef := gen.Reference()

				parcel := testutils.NewParcelMock(t).
					DefaultTargetMock.Return(&obj).
					MessageMock.Return(
					&message.AdditionalCallFromPreviousExecutor{
						ObjectReference: obj,
						RequestRef:      reqRef,
						Request:         record.IncomingRequest{},
						Pending:         insolar.NotPending,
					},
				)

				h := &HandleAdditionalCallFromPreviousExecutor{
					dep: &Dependencies{
						Sender: bus.NewSenderMock(t).ReplyMock.Return(),
						StateStorage: NewStateStorageMock(t).
							UpsertExecutionStateMock.Expect(obj).
							Return(
								NewExecutionBrokerIMock(t).
									AddAdditionalRequestFromPrevExecutorMock.Return().
									SetNotPendingMock.Return(),
							),
					},
					Parcel: parcel,
				}
				f := flow.NewFlowMock(t)
				return h, f
			},
		},
	}

	for _, test := range table {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx := inslogger.TestContext(t)
			mc := minimock.NewController(t)

			h, f := test.mocks(mc)
			err := h.Present(ctx, f)
			if test.error {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			mc.Wait(1 * time.Minute)
			mc.Finish()
		})
	}
}