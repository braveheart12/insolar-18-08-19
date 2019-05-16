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

package dispatcher

import (
	"context"
	"errors"
	"sync/atomic"

	"github.com/insolar/insolar/log"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/insolar/insolar/instrumentation/inslogger"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/flow"
	"github.com/insolar/insolar/insolar/flow/bus"
	"github.com/insolar/insolar/insolar/flow/internal/pulse"
	"github.com/insolar/insolar/insolar/flow/internal/thread"
)

const TraceIDField = "TraceID"

type Dispatcher struct {
	handles struct {
		present flow.MakeHandle
		future  flow.MakeHandle
	}
	controller         *thread.Controller
	currentPulseNumber uint32
}

func NewDispatcher(present flow.MakeHandle, future flow.MakeHandle) *Dispatcher {
	log.Debug("NEW DISPATCHER CREATED")
	d := &Dispatcher{
		controller: thread.NewController(),
	}
	d.handles.present = present
	d.handles.future = future
	d.currentPulseNumber = insolar.FirstPulseNumber
	return d
}

// ChangePulse is a handle for pulse change vent.
func (d *Dispatcher) ChangePulse(ctx context.Context, pulse insolar.Pulse) {
	log.Debug("WrapBusHandle-CHANGE PULSE ", uint32(pulse.PulseNumber))
	d.controller.Pulse()
	atomic.StoreUint32(&d.currentPulseNumber, uint32(pulse.PulseNumber)) // TODO FIXME WTF???
}

func (d *Dispatcher) getHandleByPulse(msgPulseNumber insolar.PulseNumber) flow.MakeHandle {
	log.Debug("WrapBusHandle-getHandleByPulse msgPulseNumber = ", msgPulseNumber, "current = ", atomic.LoadUint32(&d.currentPulseNumber))
	if uint32(msgPulseNumber) > atomic.LoadUint32(&d.currentPulseNumber) {
		return d.handles.future
	}
	return d.handles.present
}

func (d *Dispatcher) WrapBusHandle(ctx context.Context, parcel insolar.Parcel) (insolar.Reply, error) {
	log.Debug("WrapBusHandle-BEGIN")
	msg := bus.Message{
		ReplyTo: make(chan bus.Reply, 1),
		Parcel:  parcel,
	}

	ctx = pulse.ContextWith(ctx, parcel.Pulse())

	f := thread.NewThread(msg, d.controller)
	handle := d.getHandleByPulse(parcel.Pulse())

	log.Debugf("WrapBusHandle-BEFORE-RUN, handle = %v", handle)

	err := f.Run(ctx, handle(msg))
	log.Debug("WrapBusHandle-BEFORE-SELECT")
	var rep bus.Reply
	select {
	case rep = <-msg.ReplyTo:
		return rep.Reply, rep.Err
	default:
	}

	log.Debug("WrapBusHandle-AFTER-SELECT")

	if err != nil {
		return nil, err
	}

	log.Debug("WrapBusHandle-END")
	return nil, errors.New("no reply from handler")
}

func (d *Dispatcher) InnerSubscriber(watermillMsg *message.Message) ([]*message.Message, error) {
	msg := bus.Message{
		WatermillMsg: watermillMsg,
	}

	ctx := context.Background()
	ctx = inslogger.ContextWithTrace(ctx, watermillMsg.Metadata.Get(TraceIDField))
	logger := inslogger.FromContext(ctx)
	go func() {
		f := thread.NewThread(msg, d.controller)
		err := f.Run(ctx, d.handles.present(msg))
		if err != nil {
			logger.Error("Handling failed", err)
		}
	}()
	return nil, nil
}

// Process handles incoming message.
func (d *Dispatcher) Process(msg *message.Message) ([]*message.Message, error) {
	return nil, nil
}
