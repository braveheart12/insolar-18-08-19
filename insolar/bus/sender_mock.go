package bus

// Code generated by http://github.com/gojuno/minimock (3.0.0). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gojuno/minimock/v3"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/payload"
)

// SenderMock implements Sender
type SenderMock struct {
	t minimock.Tester

	funcLatestPulse          func(ctx context.Context) (p1 insolar.Pulse, err error)
	inspectFuncLatestPulse   func(ctx context.Context)
	afterLatestPulseCounter  uint64
	beforeLatestPulseCounter uint64
	LatestPulseMock          mSenderMockLatestPulse

	funcReply          func(ctx context.Context, origin payload.Meta, reply *message.Message)
	inspectFuncReply   func(ctx context.Context, origin payload.Meta, reply *message.Message)
	afterReplyCounter  uint64
	beforeReplyCounter uint64
	ReplyMock          mSenderMockReply

	funcSendRole          func(ctx context.Context, msg *message.Message, role insolar.DynamicRole, object insolar.Reference) (ch1 <-chan *message.Message, f1 func())
	inspectFuncSendRole   func(ctx context.Context, msg *message.Message, role insolar.DynamicRole, object insolar.Reference)
	afterSendRoleCounter  uint64
	beforeSendRoleCounter uint64
	SendRoleMock          mSenderMockSendRole

	funcSendTarget          func(ctx context.Context, msg *message.Message, target insolar.Reference) (ch1 <-chan *message.Message, f1 func())
	inspectFuncSendTarget   func(ctx context.Context, msg *message.Message, target insolar.Reference)
	afterSendTargetCounter  uint64
	beforeSendTargetCounter uint64
	SendTargetMock          mSenderMockSendTarget
}

// NewSenderMock returns a mock for Sender
func NewSenderMock(t minimock.Tester) *SenderMock {
	m := &SenderMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.LatestPulseMock = mSenderMockLatestPulse{mock: m}
	m.LatestPulseMock.callArgs = []*SenderMockLatestPulseParams{}

	m.ReplyMock = mSenderMockReply{mock: m}
	m.ReplyMock.callArgs = []*SenderMockReplyParams{}

	m.SendRoleMock = mSenderMockSendRole{mock: m}
	m.SendRoleMock.callArgs = []*SenderMockSendRoleParams{}

	m.SendTargetMock = mSenderMockSendTarget{mock: m}
	m.SendTargetMock.callArgs = []*SenderMockSendTargetParams{}

	return m
}

type mSenderMockLatestPulse struct {
	mock               *SenderMock
	defaultExpectation *SenderMockLatestPulseExpectation
	expectations       []*SenderMockLatestPulseExpectation

	callArgs []*SenderMockLatestPulseParams
	mutex    sync.RWMutex
}

// SenderMockLatestPulseExpectation specifies expectation struct of the Sender.LatestPulse
type SenderMockLatestPulseExpectation struct {
	mock    *SenderMock
	params  *SenderMockLatestPulseParams
	results *SenderMockLatestPulseResults
	Counter uint64
}

// SenderMockLatestPulseParams contains parameters of the Sender.LatestPulse
type SenderMockLatestPulseParams struct {
	ctx context.Context
}

// SenderMockLatestPulseResults contains results of the Sender.LatestPulse
type SenderMockLatestPulseResults struct {
	p1  insolar.Pulse
	err error
}

// Expect sets up expected params for Sender.LatestPulse
func (mmLatestPulse *mSenderMockLatestPulse) Expect(ctx context.Context) *mSenderMockLatestPulse {
	if mmLatestPulse.mock.funcLatestPulse != nil {
		mmLatestPulse.mock.t.Fatalf("SenderMock.LatestPulse mock is already set by Set")
	}

	if mmLatestPulse.defaultExpectation == nil {
		mmLatestPulse.defaultExpectation = &SenderMockLatestPulseExpectation{}
	}

	mmLatestPulse.defaultExpectation.params = &SenderMockLatestPulseParams{ctx}
	for _, e := range mmLatestPulse.expectations {
		if minimock.Equal(e.params, mmLatestPulse.defaultExpectation.params) {
			mmLatestPulse.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmLatestPulse.defaultExpectation.params)
		}
	}

	return mmLatestPulse
}

// Inspect accepts an inspector function that has same arguments as the Sender.LatestPulse
func (mmLatestPulse *mSenderMockLatestPulse) Inspect(f func(ctx context.Context)) *mSenderMockLatestPulse {
	if mmLatestPulse.mock.inspectFuncLatestPulse != nil {
		mmLatestPulse.mock.t.Fatalf("Inspect function is already set for SenderMock.LatestPulse")
	}

	mmLatestPulse.mock.inspectFuncLatestPulse = f

	return mmLatestPulse
}

// Return sets up results that will be returned by Sender.LatestPulse
func (mmLatestPulse *mSenderMockLatestPulse) Return(p1 insolar.Pulse, err error) *SenderMock {
	if mmLatestPulse.mock.funcLatestPulse != nil {
		mmLatestPulse.mock.t.Fatalf("SenderMock.LatestPulse mock is already set by Set")
	}

	if mmLatestPulse.defaultExpectation == nil {
		mmLatestPulse.defaultExpectation = &SenderMockLatestPulseExpectation{mock: mmLatestPulse.mock}
	}
	mmLatestPulse.defaultExpectation.results = &SenderMockLatestPulseResults{p1, err}
	return mmLatestPulse.mock
}

//Set uses given function f to mock the Sender.LatestPulse method
func (mmLatestPulse *mSenderMockLatestPulse) Set(f func(ctx context.Context) (p1 insolar.Pulse, err error)) *SenderMock {
	if mmLatestPulse.defaultExpectation != nil {
		mmLatestPulse.mock.t.Fatalf("Default expectation is already set for the Sender.LatestPulse method")
	}

	if len(mmLatestPulse.expectations) > 0 {
		mmLatestPulse.mock.t.Fatalf("Some expectations are already set for the Sender.LatestPulse method")
	}

	mmLatestPulse.mock.funcLatestPulse = f
	return mmLatestPulse.mock
}

// When sets expectation for the Sender.LatestPulse which will trigger the result defined by the following
// Then helper
func (mmLatestPulse *mSenderMockLatestPulse) When(ctx context.Context) *SenderMockLatestPulseExpectation {
	if mmLatestPulse.mock.funcLatestPulse != nil {
		mmLatestPulse.mock.t.Fatalf("SenderMock.LatestPulse mock is already set by Set")
	}

	expectation := &SenderMockLatestPulseExpectation{
		mock:   mmLatestPulse.mock,
		params: &SenderMockLatestPulseParams{ctx},
	}
	mmLatestPulse.expectations = append(mmLatestPulse.expectations, expectation)
	return expectation
}

// Then sets up Sender.LatestPulse return parameters for the expectation previously defined by the When method
func (e *SenderMockLatestPulseExpectation) Then(p1 insolar.Pulse, err error) *SenderMock {
	e.results = &SenderMockLatestPulseResults{p1, err}
	return e.mock
}

// LatestPulse implements Sender
func (mmLatestPulse *SenderMock) LatestPulse(ctx context.Context) (p1 insolar.Pulse, err error) {
	mm_atomic.AddUint64(&mmLatestPulse.beforeLatestPulseCounter, 1)
	defer mm_atomic.AddUint64(&mmLatestPulse.afterLatestPulseCounter, 1)

	if mmLatestPulse.inspectFuncLatestPulse != nil {
		mmLatestPulse.inspectFuncLatestPulse(ctx)
	}

	params := &SenderMockLatestPulseParams{ctx}

	// Record call args
	mmLatestPulse.LatestPulseMock.mutex.Lock()
	mmLatestPulse.LatestPulseMock.callArgs = append(mmLatestPulse.LatestPulseMock.callArgs, params)
	mmLatestPulse.LatestPulseMock.mutex.Unlock()

	for _, e := range mmLatestPulse.LatestPulseMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.p1, e.results.err
		}
	}

	if mmLatestPulse.LatestPulseMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmLatestPulse.LatestPulseMock.defaultExpectation.Counter, 1)
		want := mmLatestPulse.LatestPulseMock.defaultExpectation.params
		got := SenderMockLatestPulseParams{ctx}
		if want != nil && !minimock.Equal(*want, got) {
			mmLatestPulse.t.Errorf("SenderMock.LatestPulse got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmLatestPulse.LatestPulseMock.defaultExpectation.results
		if results == nil {
			mmLatestPulse.t.Fatal("No results are set for the SenderMock.LatestPulse")
		}
		return (*results).p1, (*results).err
	}
	if mmLatestPulse.funcLatestPulse != nil {
		return mmLatestPulse.funcLatestPulse(ctx)
	}
	mmLatestPulse.t.Fatalf("Unexpected call to SenderMock.LatestPulse. %v", ctx)
	return
}

// LatestPulseAfterCounter returns a count of finished SenderMock.LatestPulse invocations
func (mmLatestPulse *SenderMock) LatestPulseAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLatestPulse.afterLatestPulseCounter)
}

// LatestPulseBeforeCounter returns a count of SenderMock.LatestPulse invocations
func (mmLatestPulse *SenderMock) LatestPulseBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLatestPulse.beforeLatestPulseCounter)
}

// Calls returns a list of arguments used in each call to SenderMock.LatestPulse.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmLatestPulse *mSenderMockLatestPulse) Calls() []*SenderMockLatestPulseParams {
	mmLatestPulse.mutex.RLock()

	argCopy := make([]*SenderMockLatestPulseParams, len(mmLatestPulse.callArgs))
	copy(argCopy, mmLatestPulse.callArgs)

	mmLatestPulse.mutex.RUnlock()

	return argCopy
}

// MinimockLatestPulseDone returns true if the count of the LatestPulse invocations corresponds
// the number of defined expectations
func (m *SenderMock) MinimockLatestPulseDone() bool {
	for _, e := range m.LatestPulseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.LatestPulseMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterLatestPulseCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcLatestPulse != nil && mm_atomic.LoadUint64(&m.afterLatestPulseCounter) < 1 {
		return false
	}
	return true
}

// MinimockLatestPulseInspect logs each unmet expectation
func (m *SenderMock) MinimockLatestPulseInspect() {
	for _, e := range m.LatestPulseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to SenderMock.LatestPulse with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.LatestPulseMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterLatestPulseCounter) < 1 {
		if m.LatestPulseMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to SenderMock.LatestPulse")
		} else {
			m.t.Errorf("Expected call to SenderMock.LatestPulse with params: %#v", *m.LatestPulseMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcLatestPulse != nil && mm_atomic.LoadUint64(&m.afterLatestPulseCounter) < 1 {
		m.t.Error("Expected call to SenderMock.LatestPulse")
	}
}

type mSenderMockReply struct {
	mock               *SenderMock
	defaultExpectation *SenderMockReplyExpectation
	expectations       []*SenderMockReplyExpectation

	callArgs []*SenderMockReplyParams
	mutex    sync.RWMutex
}

// SenderMockReplyExpectation specifies expectation struct of the Sender.Reply
type SenderMockReplyExpectation struct {
	mock   *SenderMock
	params *SenderMockReplyParams

	Counter uint64
}

// SenderMockReplyParams contains parameters of the Sender.Reply
type SenderMockReplyParams struct {
	ctx    context.Context
	origin payload.Meta
	reply  *message.Message
}

// Expect sets up expected params for Sender.Reply
func (mmReply *mSenderMockReply) Expect(ctx context.Context, origin payload.Meta, reply *message.Message) *mSenderMockReply {
	if mmReply.mock.funcReply != nil {
		mmReply.mock.t.Fatalf("SenderMock.Reply mock is already set by Set")
	}

	if mmReply.defaultExpectation == nil {
		mmReply.defaultExpectation = &SenderMockReplyExpectation{}
	}

	mmReply.defaultExpectation.params = &SenderMockReplyParams{ctx, origin, reply}
	for _, e := range mmReply.expectations {
		if minimock.Equal(e.params, mmReply.defaultExpectation.params) {
			mmReply.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmReply.defaultExpectation.params)
		}
	}

	return mmReply
}

// Inspect accepts an inspector function that has same arguments as the Sender.Reply
func (mmReply *mSenderMockReply) Inspect(f func(ctx context.Context, origin payload.Meta, reply *message.Message)) *mSenderMockReply {
	if mmReply.mock.inspectFuncReply != nil {
		mmReply.mock.t.Fatalf("Inspect function is already set for SenderMock.Reply")
	}

	mmReply.mock.inspectFuncReply = f

	return mmReply
}

// Return sets up results that will be returned by Sender.Reply
func (mmReply *mSenderMockReply) Return() *SenderMock {
	if mmReply.mock.funcReply != nil {
		mmReply.mock.t.Fatalf("SenderMock.Reply mock is already set by Set")
	}

	if mmReply.defaultExpectation == nil {
		mmReply.defaultExpectation = &SenderMockReplyExpectation{mock: mmReply.mock}
	}

	return mmReply.mock
}

//Set uses given function f to mock the Sender.Reply method
func (mmReply *mSenderMockReply) Set(f func(ctx context.Context, origin payload.Meta, reply *message.Message)) *SenderMock {
	if mmReply.defaultExpectation != nil {
		mmReply.mock.t.Fatalf("Default expectation is already set for the Sender.Reply method")
	}

	if len(mmReply.expectations) > 0 {
		mmReply.mock.t.Fatalf("Some expectations are already set for the Sender.Reply method")
	}

	mmReply.mock.funcReply = f
	return mmReply.mock
}

// Reply implements Sender
func (mmReply *SenderMock) Reply(ctx context.Context, origin payload.Meta, reply *message.Message) {
	mm_atomic.AddUint64(&mmReply.beforeReplyCounter, 1)
	defer mm_atomic.AddUint64(&mmReply.afterReplyCounter, 1)

	if mmReply.inspectFuncReply != nil {
		mmReply.inspectFuncReply(ctx, origin, reply)
	}

	params := &SenderMockReplyParams{ctx, origin, reply}

	// Record call args
	mmReply.ReplyMock.mutex.Lock()
	mmReply.ReplyMock.callArgs = append(mmReply.ReplyMock.callArgs, params)
	mmReply.ReplyMock.mutex.Unlock()

	for _, e := range mmReply.ReplyMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmReply.ReplyMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmReply.ReplyMock.defaultExpectation.Counter, 1)
		want := mmReply.ReplyMock.defaultExpectation.params
		got := SenderMockReplyParams{ctx, origin, reply}
		if want != nil && !minimock.Equal(*want, got) {
			mmReply.t.Errorf("SenderMock.Reply got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		return

	}
	if mmReply.funcReply != nil {
		mmReply.funcReply(ctx, origin, reply)
		return
	}
	mmReply.t.Fatalf("Unexpected call to SenderMock.Reply. %v %v %v", ctx, origin, reply)

}

// ReplyAfterCounter returns a count of finished SenderMock.Reply invocations
func (mmReply *SenderMock) ReplyAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmReply.afterReplyCounter)
}

// ReplyBeforeCounter returns a count of SenderMock.Reply invocations
func (mmReply *SenderMock) ReplyBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmReply.beforeReplyCounter)
}

// Calls returns a list of arguments used in each call to SenderMock.Reply.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmReply *mSenderMockReply) Calls() []*SenderMockReplyParams {
	mmReply.mutex.RLock()

	argCopy := make([]*SenderMockReplyParams, len(mmReply.callArgs))
	copy(argCopy, mmReply.callArgs)

	mmReply.mutex.RUnlock()

	return argCopy
}

// MinimockReplyDone returns true if the count of the Reply invocations corresponds
// the number of defined expectations
func (m *SenderMock) MinimockReplyDone() bool {
	for _, e := range m.ReplyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ReplyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterReplyCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcReply != nil && mm_atomic.LoadUint64(&m.afterReplyCounter) < 1 {
		return false
	}
	return true
}

// MinimockReplyInspect logs each unmet expectation
func (m *SenderMock) MinimockReplyInspect() {
	for _, e := range m.ReplyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to SenderMock.Reply with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ReplyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterReplyCounter) < 1 {
		if m.ReplyMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to SenderMock.Reply")
		} else {
			m.t.Errorf("Expected call to SenderMock.Reply with params: %#v", *m.ReplyMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcReply != nil && mm_atomic.LoadUint64(&m.afterReplyCounter) < 1 {
		m.t.Error("Expected call to SenderMock.Reply")
	}
}

type mSenderMockSendRole struct {
	mock               *SenderMock
	defaultExpectation *SenderMockSendRoleExpectation
	expectations       []*SenderMockSendRoleExpectation

	callArgs []*SenderMockSendRoleParams
	mutex    sync.RWMutex
}

// SenderMockSendRoleExpectation specifies expectation struct of the Sender.SendRole
type SenderMockSendRoleExpectation struct {
	mock    *SenderMock
	params  *SenderMockSendRoleParams
	results *SenderMockSendRoleResults
	Counter uint64
}

// SenderMockSendRoleParams contains parameters of the Sender.SendRole
type SenderMockSendRoleParams struct {
	ctx    context.Context
	msg    *message.Message
	role   insolar.DynamicRole
	object insolar.Reference
}

// SenderMockSendRoleResults contains results of the Sender.SendRole
type SenderMockSendRoleResults struct {
	ch1 <-chan *message.Message
	f1  func()
}

// Expect sets up expected params for Sender.SendRole
func (mmSendRole *mSenderMockSendRole) Expect(ctx context.Context, msg *message.Message, role insolar.DynamicRole, object insolar.Reference) *mSenderMockSendRole {
	if mmSendRole.mock.funcSendRole != nil {
		mmSendRole.mock.t.Fatalf("SenderMock.SendRole mock is already set by Set")
	}

	if mmSendRole.defaultExpectation == nil {
		mmSendRole.defaultExpectation = &SenderMockSendRoleExpectation{}
	}

	mmSendRole.defaultExpectation.params = &SenderMockSendRoleParams{ctx, msg, role, object}
	for _, e := range mmSendRole.expectations {
		if minimock.Equal(e.params, mmSendRole.defaultExpectation.params) {
			mmSendRole.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSendRole.defaultExpectation.params)
		}
	}

	return mmSendRole
}

// Inspect accepts an inspector function that has same arguments as the Sender.SendRole
func (mmSendRole *mSenderMockSendRole) Inspect(f func(ctx context.Context, msg *message.Message, role insolar.DynamicRole, object insolar.Reference)) *mSenderMockSendRole {
	if mmSendRole.mock.inspectFuncSendRole != nil {
		mmSendRole.mock.t.Fatalf("Inspect function is already set for SenderMock.SendRole")
	}

	mmSendRole.mock.inspectFuncSendRole = f

	return mmSendRole
}

// Return sets up results that will be returned by Sender.SendRole
func (mmSendRole *mSenderMockSendRole) Return(ch1 <-chan *message.Message, f1 func()) *SenderMock {
	if mmSendRole.mock.funcSendRole != nil {
		mmSendRole.mock.t.Fatalf("SenderMock.SendRole mock is already set by Set")
	}

	if mmSendRole.defaultExpectation == nil {
		mmSendRole.defaultExpectation = &SenderMockSendRoleExpectation{mock: mmSendRole.mock}
	}
	mmSendRole.defaultExpectation.results = &SenderMockSendRoleResults{ch1, f1}
	return mmSendRole.mock
}

//Set uses given function f to mock the Sender.SendRole method
func (mmSendRole *mSenderMockSendRole) Set(f func(ctx context.Context, msg *message.Message, role insolar.DynamicRole, object insolar.Reference) (ch1 <-chan *message.Message, f1 func())) *SenderMock {
	if mmSendRole.defaultExpectation != nil {
		mmSendRole.mock.t.Fatalf("Default expectation is already set for the Sender.SendRole method")
	}

	if len(mmSendRole.expectations) > 0 {
		mmSendRole.mock.t.Fatalf("Some expectations are already set for the Sender.SendRole method")
	}

	mmSendRole.mock.funcSendRole = f
	return mmSendRole.mock
}

// When sets expectation for the Sender.SendRole which will trigger the result defined by the following
// Then helper
func (mmSendRole *mSenderMockSendRole) When(ctx context.Context, msg *message.Message, role insolar.DynamicRole, object insolar.Reference) *SenderMockSendRoleExpectation {
	if mmSendRole.mock.funcSendRole != nil {
		mmSendRole.mock.t.Fatalf("SenderMock.SendRole mock is already set by Set")
	}

	expectation := &SenderMockSendRoleExpectation{
		mock:   mmSendRole.mock,
		params: &SenderMockSendRoleParams{ctx, msg, role, object},
	}
	mmSendRole.expectations = append(mmSendRole.expectations, expectation)
	return expectation
}

// Then sets up Sender.SendRole return parameters for the expectation previously defined by the When method
func (e *SenderMockSendRoleExpectation) Then(ch1 <-chan *message.Message, f1 func()) *SenderMock {
	e.results = &SenderMockSendRoleResults{ch1, f1}
	return e.mock
}

// SendRole implements Sender
func (mmSendRole *SenderMock) SendRole(ctx context.Context, msg *message.Message, role insolar.DynamicRole, object insolar.Reference) (ch1 <-chan *message.Message, f1 func()) {
	mm_atomic.AddUint64(&mmSendRole.beforeSendRoleCounter, 1)
	defer mm_atomic.AddUint64(&mmSendRole.afterSendRoleCounter, 1)

	if mmSendRole.inspectFuncSendRole != nil {
		mmSendRole.inspectFuncSendRole(ctx, msg, role, object)
	}

	params := &SenderMockSendRoleParams{ctx, msg, role, object}

	// Record call args
	mmSendRole.SendRoleMock.mutex.Lock()
	mmSendRole.SendRoleMock.callArgs = append(mmSendRole.SendRoleMock.callArgs, params)
	mmSendRole.SendRoleMock.mutex.Unlock()

	for _, e := range mmSendRole.SendRoleMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ch1, e.results.f1
		}
	}

	if mmSendRole.SendRoleMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSendRole.SendRoleMock.defaultExpectation.Counter, 1)
		want := mmSendRole.SendRoleMock.defaultExpectation.params
		got := SenderMockSendRoleParams{ctx, msg, role, object}
		if want != nil && !minimock.Equal(*want, got) {
			mmSendRole.t.Errorf("SenderMock.SendRole got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmSendRole.SendRoleMock.defaultExpectation.results
		if results == nil {
			mmSendRole.t.Fatal("No results are set for the SenderMock.SendRole")
		}
		return (*results).ch1, (*results).f1
	}
	if mmSendRole.funcSendRole != nil {
		return mmSendRole.funcSendRole(ctx, msg, role, object)
	}
	mmSendRole.t.Fatalf("Unexpected call to SenderMock.SendRole. %v %v %v %v", ctx, msg, role, object)
	return
}

// SendRoleAfterCounter returns a count of finished SenderMock.SendRole invocations
func (mmSendRole *SenderMock) SendRoleAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendRole.afterSendRoleCounter)
}

// SendRoleBeforeCounter returns a count of SenderMock.SendRole invocations
func (mmSendRole *SenderMock) SendRoleBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendRole.beforeSendRoleCounter)
}

// Calls returns a list of arguments used in each call to SenderMock.SendRole.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSendRole *mSenderMockSendRole) Calls() []*SenderMockSendRoleParams {
	mmSendRole.mutex.RLock()

	argCopy := make([]*SenderMockSendRoleParams, len(mmSendRole.callArgs))
	copy(argCopy, mmSendRole.callArgs)

	mmSendRole.mutex.RUnlock()

	return argCopy
}

// MinimockSendRoleDone returns true if the count of the SendRole invocations corresponds
// the number of defined expectations
func (m *SenderMock) MinimockSendRoleDone() bool {
	for _, e := range m.SendRoleMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendRoleMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendRoleCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendRole != nil && mm_atomic.LoadUint64(&m.afterSendRoleCounter) < 1 {
		return false
	}
	return true
}

// MinimockSendRoleInspect logs each unmet expectation
func (m *SenderMock) MinimockSendRoleInspect() {
	for _, e := range m.SendRoleMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to SenderMock.SendRole with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendRoleMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendRoleCounter) < 1 {
		if m.SendRoleMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to SenderMock.SendRole")
		} else {
			m.t.Errorf("Expected call to SenderMock.SendRole with params: %#v", *m.SendRoleMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendRole != nil && mm_atomic.LoadUint64(&m.afterSendRoleCounter) < 1 {
		m.t.Error("Expected call to SenderMock.SendRole")
	}
}

type mSenderMockSendTarget struct {
	mock               *SenderMock
	defaultExpectation *SenderMockSendTargetExpectation
	expectations       []*SenderMockSendTargetExpectation

	callArgs []*SenderMockSendTargetParams
	mutex    sync.RWMutex
}

// SenderMockSendTargetExpectation specifies expectation struct of the Sender.SendTarget
type SenderMockSendTargetExpectation struct {
	mock    *SenderMock
	params  *SenderMockSendTargetParams
	results *SenderMockSendTargetResults
	Counter uint64
}

// SenderMockSendTargetParams contains parameters of the Sender.SendTarget
type SenderMockSendTargetParams struct {
	ctx    context.Context
	msg    *message.Message
	target insolar.Reference
}

// SenderMockSendTargetResults contains results of the Sender.SendTarget
type SenderMockSendTargetResults struct {
	ch1 <-chan *message.Message
	f1  func()
}

// Expect sets up expected params for Sender.SendTarget
func (mmSendTarget *mSenderMockSendTarget) Expect(ctx context.Context, msg *message.Message, target insolar.Reference) *mSenderMockSendTarget {
	if mmSendTarget.mock.funcSendTarget != nil {
		mmSendTarget.mock.t.Fatalf("SenderMock.SendTarget mock is already set by Set")
	}

	if mmSendTarget.defaultExpectation == nil {
		mmSendTarget.defaultExpectation = &SenderMockSendTargetExpectation{}
	}

	mmSendTarget.defaultExpectation.params = &SenderMockSendTargetParams{ctx, msg, target}
	for _, e := range mmSendTarget.expectations {
		if minimock.Equal(e.params, mmSendTarget.defaultExpectation.params) {
			mmSendTarget.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSendTarget.defaultExpectation.params)
		}
	}

	return mmSendTarget
}

// Inspect accepts an inspector function that has same arguments as the Sender.SendTarget
func (mmSendTarget *mSenderMockSendTarget) Inspect(f func(ctx context.Context, msg *message.Message, target insolar.Reference)) *mSenderMockSendTarget {
	if mmSendTarget.mock.inspectFuncSendTarget != nil {
		mmSendTarget.mock.t.Fatalf("Inspect function is already set for SenderMock.SendTarget")
	}

	mmSendTarget.mock.inspectFuncSendTarget = f

	return mmSendTarget
}

// Return sets up results that will be returned by Sender.SendTarget
func (mmSendTarget *mSenderMockSendTarget) Return(ch1 <-chan *message.Message, f1 func()) *SenderMock {
	if mmSendTarget.mock.funcSendTarget != nil {
		mmSendTarget.mock.t.Fatalf("SenderMock.SendTarget mock is already set by Set")
	}

	if mmSendTarget.defaultExpectation == nil {
		mmSendTarget.defaultExpectation = &SenderMockSendTargetExpectation{mock: mmSendTarget.mock}
	}
	mmSendTarget.defaultExpectation.results = &SenderMockSendTargetResults{ch1, f1}
	return mmSendTarget.mock
}

//Set uses given function f to mock the Sender.SendTarget method
func (mmSendTarget *mSenderMockSendTarget) Set(f func(ctx context.Context, msg *message.Message, target insolar.Reference) (ch1 <-chan *message.Message, f1 func())) *SenderMock {
	if mmSendTarget.defaultExpectation != nil {
		mmSendTarget.mock.t.Fatalf("Default expectation is already set for the Sender.SendTarget method")
	}

	if len(mmSendTarget.expectations) > 0 {
		mmSendTarget.mock.t.Fatalf("Some expectations are already set for the Sender.SendTarget method")
	}

	mmSendTarget.mock.funcSendTarget = f
	return mmSendTarget.mock
}

// When sets expectation for the Sender.SendTarget which will trigger the result defined by the following
// Then helper
func (mmSendTarget *mSenderMockSendTarget) When(ctx context.Context, msg *message.Message, target insolar.Reference) *SenderMockSendTargetExpectation {
	if mmSendTarget.mock.funcSendTarget != nil {
		mmSendTarget.mock.t.Fatalf("SenderMock.SendTarget mock is already set by Set")
	}

	expectation := &SenderMockSendTargetExpectation{
		mock:   mmSendTarget.mock,
		params: &SenderMockSendTargetParams{ctx, msg, target},
	}
	mmSendTarget.expectations = append(mmSendTarget.expectations, expectation)
	return expectation
}

// Then sets up Sender.SendTarget return parameters for the expectation previously defined by the When method
func (e *SenderMockSendTargetExpectation) Then(ch1 <-chan *message.Message, f1 func()) *SenderMock {
	e.results = &SenderMockSendTargetResults{ch1, f1}
	return e.mock
}

// SendTarget implements Sender
func (mmSendTarget *SenderMock) SendTarget(ctx context.Context, msg *message.Message, target insolar.Reference) (ch1 <-chan *message.Message, f1 func()) {
	mm_atomic.AddUint64(&mmSendTarget.beforeSendTargetCounter, 1)
	defer mm_atomic.AddUint64(&mmSendTarget.afterSendTargetCounter, 1)

	if mmSendTarget.inspectFuncSendTarget != nil {
		mmSendTarget.inspectFuncSendTarget(ctx, msg, target)
	}

	params := &SenderMockSendTargetParams{ctx, msg, target}

	// Record call args
	mmSendTarget.SendTargetMock.mutex.Lock()
	mmSendTarget.SendTargetMock.callArgs = append(mmSendTarget.SendTargetMock.callArgs, params)
	mmSendTarget.SendTargetMock.mutex.Unlock()

	for _, e := range mmSendTarget.SendTargetMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ch1, e.results.f1
		}
	}

	if mmSendTarget.SendTargetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSendTarget.SendTargetMock.defaultExpectation.Counter, 1)
		want := mmSendTarget.SendTargetMock.defaultExpectation.params
		got := SenderMockSendTargetParams{ctx, msg, target}
		if want != nil && !minimock.Equal(*want, got) {
			mmSendTarget.t.Errorf("SenderMock.SendTarget got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmSendTarget.SendTargetMock.defaultExpectation.results
		if results == nil {
			mmSendTarget.t.Fatal("No results are set for the SenderMock.SendTarget")
		}
		return (*results).ch1, (*results).f1
	}
	if mmSendTarget.funcSendTarget != nil {
		return mmSendTarget.funcSendTarget(ctx, msg, target)
	}
	mmSendTarget.t.Fatalf("Unexpected call to SenderMock.SendTarget. %v %v %v", ctx, msg, target)
	return
}

// SendTargetAfterCounter returns a count of finished SenderMock.SendTarget invocations
func (mmSendTarget *SenderMock) SendTargetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendTarget.afterSendTargetCounter)
}

// SendTargetBeforeCounter returns a count of SenderMock.SendTarget invocations
func (mmSendTarget *SenderMock) SendTargetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendTarget.beforeSendTargetCounter)
}

// Calls returns a list of arguments used in each call to SenderMock.SendTarget.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSendTarget *mSenderMockSendTarget) Calls() []*SenderMockSendTargetParams {
	mmSendTarget.mutex.RLock()

	argCopy := make([]*SenderMockSendTargetParams, len(mmSendTarget.callArgs))
	copy(argCopy, mmSendTarget.callArgs)

	mmSendTarget.mutex.RUnlock()

	return argCopy
}

// MinimockSendTargetDone returns true if the count of the SendTarget invocations corresponds
// the number of defined expectations
func (m *SenderMock) MinimockSendTargetDone() bool {
	for _, e := range m.SendTargetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendTargetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendTargetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendTarget != nil && mm_atomic.LoadUint64(&m.afterSendTargetCounter) < 1 {
		return false
	}
	return true
}

// MinimockSendTargetInspect logs each unmet expectation
func (m *SenderMock) MinimockSendTargetInspect() {
	for _, e := range m.SendTargetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to SenderMock.SendTarget with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendTargetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendTargetCounter) < 1 {
		if m.SendTargetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to SenderMock.SendTarget")
		} else {
			m.t.Errorf("Expected call to SenderMock.SendTarget with params: %#v", *m.SendTargetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendTarget != nil && mm_atomic.LoadUint64(&m.afterSendTargetCounter) < 1 {
		m.t.Error("Expected call to SenderMock.SendTarget")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *SenderMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockLatestPulseInspect()

		m.MinimockReplyInspect()

		m.MinimockSendRoleInspect()

		m.MinimockSendTargetInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *SenderMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *SenderMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockLatestPulseDone() &&
		m.MinimockReplyDone() &&
		m.MinimockSendRoleDone() &&
		m.MinimockSendTargetDone()
}
