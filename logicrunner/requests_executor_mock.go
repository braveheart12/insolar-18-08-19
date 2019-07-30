package logicrunner

// Code generated by http://github.com/gojuno/minimock (3.0.0). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/logicrunner/artifacts"
)

// RequestsExecutorMock implements RequestsExecutor
type RequestsExecutorMock struct {
	t minimock.Tester

	funcExecute          func(ctx context.Context, current *Transcript) (r1 artifacts.RequestResult, err error)
	inspectFuncExecute   func(ctx context.Context, current *Transcript)
	afterExecuteCounter  uint64
	beforeExecuteCounter uint64
	ExecuteMock          mRequestsExecutorMockExecute

	funcExecuteAndSave          func(ctx context.Context, current *Transcript) (r1 insolar.Reply, err error)
	inspectFuncExecuteAndSave   func(ctx context.Context, current *Transcript)
	afterExecuteAndSaveCounter  uint64
	beforeExecuteAndSaveCounter uint64
	ExecuteAndSaveMock          mRequestsExecutorMockExecuteAndSave

	funcSave          func(ctx context.Context, current *Transcript, res artifacts.RequestResult) (r1 insolar.Reply, err error)
	inspectFuncSave   func(ctx context.Context, current *Transcript, res artifacts.RequestResult)
	afterSaveCounter  uint64
	beforeSaveCounter uint64
	SaveMock          mRequestsExecutorMockSave

	funcSendReply          func(ctx context.Context, current *Transcript, re insolar.Reply, err error)
	inspectFuncSendReply   func(ctx context.Context, current *Transcript, re insolar.Reply, err error)
	afterSendReplyCounter  uint64
	beforeSendReplyCounter uint64
	SendReplyMock          mRequestsExecutorMockSendReply
}

// NewRequestsExecutorMock returns a mock for RequestsExecutor
func NewRequestsExecutorMock(t minimock.Tester) *RequestsExecutorMock {
	m := &RequestsExecutorMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ExecuteMock = mRequestsExecutorMockExecute{mock: m}
	m.ExecuteMock.callArgs = []*RequestsExecutorMockExecuteParams{}

	m.ExecuteAndSaveMock = mRequestsExecutorMockExecuteAndSave{mock: m}
	m.ExecuteAndSaveMock.callArgs = []*RequestsExecutorMockExecuteAndSaveParams{}

	m.SaveMock = mRequestsExecutorMockSave{mock: m}
	m.SaveMock.callArgs = []*RequestsExecutorMockSaveParams{}

	m.SendReplyMock = mRequestsExecutorMockSendReply{mock: m}
	m.SendReplyMock.callArgs = []*RequestsExecutorMockSendReplyParams{}

	return m
}

type mRequestsExecutorMockExecute struct {
	mock               *RequestsExecutorMock
	defaultExpectation *RequestsExecutorMockExecuteExpectation
	expectations       []*RequestsExecutorMockExecuteExpectation

	callArgs []*RequestsExecutorMockExecuteParams
	mutex    sync.RWMutex
}

// RequestsExecutorMockExecuteExpectation specifies expectation struct of the RequestsExecutor.Execute
type RequestsExecutorMockExecuteExpectation struct {
	mock    *RequestsExecutorMock
	params  *RequestsExecutorMockExecuteParams
	results *RequestsExecutorMockExecuteResults
	Counter uint64
}

// RequestsExecutorMockExecuteParams contains parameters of the RequestsExecutor.Execute
type RequestsExecutorMockExecuteParams struct {
	ctx     context.Context
	current *Transcript
}

// RequestsExecutorMockExecuteResults contains results of the RequestsExecutor.Execute
type RequestsExecutorMockExecuteResults struct {
	r1  artifacts.RequestResult
	err error
}

// Expect sets up expected params for RequestsExecutor.Execute
func (mmExecute *mRequestsExecutorMockExecute) Expect(ctx context.Context, current *Transcript) *mRequestsExecutorMockExecute {
	if mmExecute.mock.funcExecute != nil {
		mmExecute.mock.t.Fatalf("RequestsExecutorMock.Execute mock is already set by Set")
	}

	if mmExecute.defaultExpectation == nil {
		mmExecute.defaultExpectation = &RequestsExecutorMockExecuteExpectation{}
	}

	mmExecute.defaultExpectation.params = &RequestsExecutorMockExecuteParams{ctx, current}
	for _, e := range mmExecute.expectations {
		if minimock.Equal(e.params, mmExecute.defaultExpectation.params) {
			mmExecute.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmExecute.defaultExpectation.params)
		}
	}

	return mmExecute
}

// Inspect accepts an inspector function that has same arguments as the RequestsExecutor.Execute
func (mmExecute *mRequestsExecutorMockExecute) Inspect(f func(ctx context.Context, current *Transcript)) *mRequestsExecutorMockExecute {
	if mmExecute.mock.inspectFuncExecute != nil {
		mmExecute.mock.t.Fatalf("Inspect function is already set for RequestsExecutorMock.Execute")
	}

	mmExecute.mock.inspectFuncExecute = f

	return mmExecute
}

// Return sets up results that will be returned by RequestsExecutor.Execute
func (mmExecute *mRequestsExecutorMockExecute) Return(r1 artifacts.RequestResult, err error) *RequestsExecutorMock {
	if mmExecute.mock.funcExecute != nil {
		mmExecute.mock.t.Fatalf("RequestsExecutorMock.Execute mock is already set by Set")
	}

	if mmExecute.defaultExpectation == nil {
		mmExecute.defaultExpectation = &RequestsExecutorMockExecuteExpectation{mock: mmExecute.mock}
	}
	mmExecute.defaultExpectation.results = &RequestsExecutorMockExecuteResults{r1, err}
	return mmExecute.mock
}

//Set uses given function f to mock the RequestsExecutor.Execute method
func (mmExecute *mRequestsExecutorMockExecute) Set(f func(ctx context.Context, current *Transcript) (r1 artifacts.RequestResult, err error)) *RequestsExecutorMock {
	if mmExecute.defaultExpectation != nil {
		mmExecute.mock.t.Fatalf("Default expectation is already set for the RequestsExecutor.Execute method")
	}

	if len(mmExecute.expectations) > 0 {
		mmExecute.mock.t.Fatalf("Some expectations are already set for the RequestsExecutor.Execute method")
	}

	mmExecute.mock.funcExecute = f
	return mmExecute.mock
}

// When sets expectation for the RequestsExecutor.Execute which will trigger the result defined by the following
// Then helper
func (mmExecute *mRequestsExecutorMockExecute) When(ctx context.Context, current *Transcript) *RequestsExecutorMockExecuteExpectation {
	if mmExecute.mock.funcExecute != nil {
		mmExecute.mock.t.Fatalf("RequestsExecutorMock.Execute mock is already set by Set")
	}

	expectation := &RequestsExecutorMockExecuteExpectation{
		mock:   mmExecute.mock,
		params: &RequestsExecutorMockExecuteParams{ctx, current},
	}
	mmExecute.expectations = append(mmExecute.expectations, expectation)
	return expectation
}

// Then sets up RequestsExecutor.Execute return parameters for the expectation previously defined by the When method
func (e *RequestsExecutorMockExecuteExpectation) Then(r1 artifacts.RequestResult, err error) *RequestsExecutorMock {
	e.results = &RequestsExecutorMockExecuteResults{r1, err}
	return e.mock
}

// Execute implements RequestsExecutor
func (mmExecute *RequestsExecutorMock) Execute(ctx context.Context, current *Transcript) (r1 artifacts.RequestResult, err error) {
	mm_atomic.AddUint64(&mmExecute.beforeExecuteCounter, 1)
	defer mm_atomic.AddUint64(&mmExecute.afterExecuteCounter, 1)

	if mmExecute.inspectFuncExecute != nil {
		mmExecute.inspectFuncExecute(ctx, current)
	}

	params := &RequestsExecutorMockExecuteParams{ctx, current}

	// Record call args
	mmExecute.ExecuteMock.mutex.Lock()
	mmExecute.ExecuteMock.callArgs = append(mmExecute.ExecuteMock.callArgs, params)
	mmExecute.ExecuteMock.mutex.Unlock()

	for _, e := range mmExecute.ExecuteMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.r1, e.results.err
		}
	}

	if mmExecute.ExecuteMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmExecute.ExecuteMock.defaultExpectation.Counter, 1)
		want := mmExecute.ExecuteMock.defaultExpectation.params
		got := RequestsExecutorMockExecuteParams{ctx, current}
		if want != nil && !minimock.Equal(*want, got) {
			mmExecute.t.Errorf("RequestsExecutorMock.Execute got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmExecute.ExecuteMock.defaultExpectation.results
		if results == nil {
			mmExecute.t.Fatal("No results are set for the RequestsExecutorMock.Execute")
		}
		return (*results).r1, (*results).err
	}
	if mmExecute.funcExecute != nil {
		return mmExecute.funcExecute(ctx, current)
	}
	mmExecute.t.Fatalf("Unexpected call to RequestsExecutorMock.Execute. %v %v", ctx, current)
	return
}

// ExecuteAfterCounter returns a count of finished RequestsExecutorMock.Execute invocations
func (mmExecute *RequestsExecutorMock) ExecuteAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmExecute.afterExecuteCounter)
}

// ExecuteBeforeCounter returns a count of RequestsExecutorMock.Execute invocations
func (mmExecute *RequestsExecutorMock) ExecuteBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmExecute.beforeExecuteCounter)
}

// Calls returns a list of arguments used in each call to RequestsExecutorMock.Execute.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmExecute *mRequestsExecutorMockExecute) Calls() []*RequestsExecutorMockExecuteParams {
	mmExecute.mutex.RLock()

	argCopy := make([]*RequestsExecutorMockExecuteParams, len(mmExecute.callArgs))
	copy(argCopy, mmExecute.callArgs)

	mmExecute.mutex.RUnlock()

	return argCopy
}

// MinimockExecuteDone returns true if the count of the Execute invocations corresponds
// the number of defined expectations
func (m *RequestsExecutorMock) MinimockExecuteDone() bool {
	for _, e := range m.ExecuteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ExecuteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterExecuteCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcExecute != nil && mm_atomic.LoadUint64(&m.afterExecuteCounter) < 1 {
		return false
	}
	return true
}

// MinimockExecuteInspect logs each unmet expectation
func (m *RequestsExecutorMock) MinimockExecuteInspect() {
	for _, e := range m.ExecuteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RequestsExecutorMock.Execute with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ExecuteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterExecuteCounter) < 1 {
		if m.ExecuteMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RequestsExecutorMock.Execute")
		} else {
			m.t.Errorf("Expected call to RequestsExecutorMock.Execute with params: %#v", *m.ExecuteMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcExecute != nil && mm_atomic.LoadUint64(&m.afterExecuteCounter) < 1 {
		m.t.Error("Expected call to RequestsExecutorMock.Execute")
	}
}

type mRequestsExecutorMockExecuteAndSave struct {
	mock               *RequestsExecutorMock
	defaultExpectation *RequestsExecutorMockExecuteAndSaveExpectation
	expectations       []*RequestsExecutorMockExecuteAndSaveExpectation

	callArgs []*RequestsExecutorMockExecuteAndSaveParams
	mutex    sync.RWMutex
}

// RequestsExecutorMockExecuteAndSaveExpectation specifies expectation struct of the RequestsExecutor.ExecuteAndSave
type RequestsExecutorMockExecuteAndSaveExpectation struct {
	mock    *RequestsExecutorMock
	params  *RequestsExecutorMockExecuteAndSaveParams
	results *RequestsExecutorMockExecuteAndSaveResults
	Counter uint64
}

// RequestsExecutorMockExecuteAndSaveParams contains parameters of the RequestsExecutor.ExecuteAndSave
type RequestsExecutorMockExecuteAndSaveParams struct {
	ctx     context.Context
	current *Transcript
}

// RequestsExecutorMockExecuteAndSaveResults contains results of the RequestsExecutor.ExecuteAndSave
type RequestsExecutorMockExecuteAndSaveResults struct {
	r1  insolar.Reply
	err error
}

// Expect sets up expected params for RequestsExecutor.ExecuteAndSave
func (mmExecuteAndSave *mRequestsExecutorMockExecuteAndSave) Expect(ctx context.Context, current *Transcript) *mRequestsExecutorMockExecuteAndSave {
	if mmExecuteAndSave.mock.funcExecuteAndSave != nil {
		mmExecuteAndSave.mock.t.Fatalf("RequestsExecutorMock.ExecuteAndSave mock is already set by Set")
	}

	if mmExecuteAndSave.defaultExpectation == nil {
		mmExecuteAndSave.defaultExpectation = &RequestsExecutorMockExecuteAndSaveExpectation{}
	}

	mmExecuteAndSave.defaultExpectation.params = &RequestsExecutorMockExecuteAndSaveParams{ctx, current}
	for _, e := range mmExecuteAndSave.expectations {
		if minimock.Equal(e.params, mmExecuteAndSave.defaultExpectation.params) {
			mmExecuteAndSave.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmExecuteAndSave.defaultExpectation.params)
		}
	}

	return mmExecuteAndSave
}

// Inspect accepts an inspector function that has same arguments as the RequestsExecutor.ExecuteAndSave
func (mmExecuteAndSave *mRequestsExecutorMockExecuteAndSave) Inspect(f func(ctx context.Context, current *Transcript)) *mRequestsExecutorMockExecuteAndSave {
	if mmExecuteAndSave.mock.inspectFuncExecuteAndSave != nil {
		mmExecuteAndSave.mock.t.Fatalf("Inspect function is already set for RequestsExecutorMock.ExecuteAndSave")
	}

	mmExecuteAndSave.mock.inspectFuncExecuteAndSave = f

	return mmExecuteAndSave
}

// Return sets up results that will be returned by RequestsExecutor.ExecuteAndSave
func (mmExecuteAndSave *mRequestsExecutorMockExecuteAndSave) Return(r1 insolar.Reply, err error) *RequestsExecutorMock {
	if mmExecuteAndSave.mock.funcExecuteAndSave != nil {
		mmExecuteAndSave.mock.t.Fatalf("RequestsExecutorMock.ExecuteAndSave mock is already set by Set")
	}

	if mmExecuteAndSave.defaultExpectation == nil {
		mmExecuteAndSave.defaultExpectation = &RequestsExecutorMockExecuteAndSaveExpectation{mock: mmExecuteAndSave.mock}
	}
	mmExecuteAndSave.defaultExpectation.results = &RequestsExecutorMockExecuteAndSaveResults{r1, err}
	return mmExecuteAndSave.mock
}

//Set uses given function f to mock the RequestsExecutor.ExecuteAndSave method
func (mmExecuteAndSave *mRequestsExecutorMockExecuteAndSave) Set(f func(ctx context.Context, current *Transcript) (r1 insolar.Reply, err error)) *RequestsExecutorMock {
	if mmExecuteAndSave.defaultExpectation != nil {
		mmExecuteAndSave.mock.t.Fatalf("Default expectation is already set for the RequestsExecutor.ExecuteAndSave method")
	}

	if len(mmExecuteAndSave.expectations) > 0 {
		mmExecuteAndSave.mock.t.Fatalf("Some expectations are already set for the RequestsExecutor.ExecuteAndSave method")
	}

	mmExecuteAndSave.mock.funcExecuteAndSave = f
	return mmExecuteAndSave.mock
}

// When sets expectation for the RequestsExecutor.ExecuteAndSave which will trigger the result defined by the following
// Then helper
func (mmExecuteAndSave *mRequestsExecutorMockExecuteAndSave) When(ctx context.Context, current *Transcript) *RequestsExecutorMockExecuteAndSaveExpectation {
	if mmExecuteAndSave.mock.funcExecuteAndSave != nil {
		mmExecuteAndSave.mock.t.Fatalf("RequestsExecutorMock.ExecuteAndSave mock is already set by Set")
	}

	expectation := &RequestsExecutorMockExecuteAndSaveExpectation{
		mock:   mmExecuteAndSave.mock,
		params: &RequestsExecutorMockExecuteAndSaveParams{ctx, current},
	}
	mmExecuteAndSave.expectations = append(mmExecuteAndSave.expectations, expectation)
	return expectation
}

// Then sets up RequestsExecutor.ExecuteAndSave return parameters for the expectation previously defined by the When method
func (e *RequestsExecutorMockExecuteAndSaveExpectation) Then(r1 insolar.Reply, err error) *RequestsExecutorMock {
	e.results = &RequestsExecutorMockExecuteAndSaveResults{r1, err}
	return e.mock
}

// ExecuteAndSave implements RequestsExecutor
func (mmExecuteAndSave *RequestsExecutorMock) ExecuteAndSave(ctx context.Context, current *Transcript) (r1 insolar.Reply, err error) {
	mm_atomic.AddUint64(&mmExecuteAndSave.beforeExecuteAndSaveCounter, 1)
	defer mm_atomic.AddUint64(&mmExecuteAndSave.afterExecuteAndSaveCounter, 1)

	if mmExecuteAndSave.inspectFuncExecuteAndSave != nil {
		mmExecuteAndSave.inspectFuncExecuteAndSave(ctx, current)
	}

	params := &RequestsExecutorMockExecuteAndSaveParams{ctx, current}

	// Record call args
	mmExecuteAndSave.ExecuteAndSaveMock.mutex.Lock()
	mmExecuteAndSave.ExecuteAndSaveMock.callArgs = append(mmExecuteAndSave.ExecuteAndSaveMock.callArgs, params)
	mmExecuteAndSave.ExecuteAndSaveMock.mutex.Unlock()

	for _, e := range mmExecuteAndSave.ExecuteAndSaveMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.r1, e.results.err
		}
	}

	if mmExecuteAndSave.ExecuteAndSaveMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmExecuteAndSave.ExecuteAndSaveMock.defaultExpectation.Counter, 1)
		want := mmExecuteAndSave.ExecuteAndSaveMock.defaultExpectation.params
		got := RequestsExecutorMockExecuteAndSaveParams{ctx, current}
		if want != nil && !minimock.Equal(*want, got) {
			mmExecuteAndSave.t.Errorf("RequestsExecutorMock.ExecuteAndSave got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmExecuteAndSave.ExecuteAndSaveMock.defaultExpectation.results
		if results == nil {
			mmExecuteAndSave.t.Fatal("No results are set for the RequestsExecutorMock.ExecuteAndSave")
		}
		return (*results).r1, (*results).err
	}
	if mmExecuteAndSave.funcExecuteAndSave != nil {
		return mmExecuteAndSave.funcExecuteAndSave(ctx, current)
	}
	mmExecuteAndSave.t.Fatalf("Unexpected call to RequestsExecutorMock.ExecuteAndSave. %v %v", ctx, current)
	return
}

// ExecuteAndSaveAfterCounter returns a count of finished RequestsExecutorMock.ExecuteAndSave invocations
func (mmExecuteAndSave *RequestsExecutorMock) ExecuteAndSaveAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmExecuteAndSave.afterExecuteAndSaveCounter)
}

// ExecuteAndSaveBeforeCounter returns a count of RequestsExecutorMock.ExecuteAndSave invocations
func (mmExecuteAndSave *RequestsExecutorMock) ExecuteAndSaveBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmExecuteAndSave.beforeExecuteAndSaveCounter)
}

// Calls returns a list of arguments used in each call to RequestsExecutorMock.ExecuteAndSave.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmExecuteAndSave *mRequestsExecutorMockExecuteAndSave) Calls() []*RequestsExecutorMockExecuteAndSaveParams {
	mmExecuteAndSave.mutex.RLock()

	argCopy := make([]*RequestsExecutorMockExecuteAndSaveParams, len(mmExecuteAndSave.callArgs))
	copy(argCopy, mmExecuteAndSave.callArgs)

	mmExecuteAndSave.mutex.RUnlock()

	return argCopy
}

// MinimockExecuteAndSaveDone returns true if the count of the ExecuteAndSave invocations corresponds
// the number of defined expectations
func (m *RequestsExecutorMock) MinimockExecuteAndSaveDone() bool {
	for _, e := range m.ExecuteAndSaveMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ExecuteAndSaveMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterExecuteAndSaveCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcExecuteAndSave != nil && mm_atomic.LoadUint64(&m.afterExecuteAndSaveCounter) < 1 {
		return false
	}
	return true
}

// MinimockExecuteAndSaveInspect logs each unmet expectation
func (m *RequestsExecutorMock) MinimockExecuteAndSaveInspect() {
	for _, e := range m.ExecuteAndSaveMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RequestsExecutorMock.ExecuteAndSave with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ExecuteAndSaveMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterExecuteAndSaveCounter) < 1 {
		if m.ExecuteAndSaveMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RequestsExecutorMock.ExecuteAndSave")
		} else {
			m.t.Errorf("Expected call to RequestsExecutorMock.ExecuteAndSave with params: %#v", *m.ExecuteAndSaveMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcExecuteAndSave != nil && mm_atomic.LoadUint64(&m.afterExecuteAndSaveCounter) < 1 {
		m.t.Error("Expected call to RequestsExecutorMock.ExecuteAndSave")
	}
}

type mRequestsExecutorMockSave struct {
	mock               *RequestsExecutorMock
	defaultExpectation *RequestsExecutorMockSaveExpectation
	expectations       []*RequestsExecutorMockSaveExpectation

	callArgs []*RequestsExecutorMockSaveParams
	mutex    sync.RWMutex
}

// RequestsExecutorMockSaveExpectation specifies expectation struct of the RequestsExecutor.Save
type RequestsExecutorMockSaveExpectation struct {
	mock    *RequestsExecutorMock
	params  *RequestsExecutorMockSaveParams
	results *RequestsExecutorMockSaveResults
	Counter uint64
}

// RequestsExecutorMockSaveParams contains parameters of the RequestsExecutor.Save
type RequestsExecutorMockSaveParams struct {
	ctx     context.Context
	current *Transcript
	res     artifacts.RequestResult
}

// RequestsExecutorMockSaveResults contains results of the RequestsExecutor.Save
type RequestsExecutorMockSaveResults struct {
	r1  insolar.Reply
	err error
}

// Expect sets up expected params for RequestsExecutor.Save
func (mmSave *mRequestsExecutorMockSave) Expect(ctx context.Context, current *Transcript, res artifacts.RequestResult) *mRequestsExecutorMockSave {
	if mmSave.mock.funcSave != nil {
		mmSave.mock.t.Fatalf("RequestsExecutorMock.Save mock is already set by Set")
	}

	if mmSave.defaultExpectation == nil {
		mmSave.defaultExpectation = &RequestsExecutorMockSaveExpectation{}
	}

	mmSave.defaultExpectation.params = &RequestsExecutorMockSaveParams{ctx, current, res}
	for _, e := range mmSave.expectations {
		if minimock.Equal(e.params, mmSave.defaultExpectation.params) {
			mmSave.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSave.defaultExpectation.params)
		}
	}

	return mmSave
}

// Inspect accepts an inspector function that has same arguments as the RequestsExecutor.Save
func (mmSave *mRequestsExecutorMockSave) Inspect(f func(ctx context.Context, current *Transcript, res artifacts.RequestResult)) *mRequestsExecutorMockSave {
	if mmSave.mock.inspectFuncSave != nil {
		mmSave.mock.t.Fatalf("Inspect function is already set for RequestsExecutorMock.Save")
	}

	mmSave.mock.inspectFuncSave = f

	return mmSave
}

// Return sets up results that will be returned by RequestsExecutor.Save
func (mmSave *mRequestsExecutorMockSave) Return(r1 insolar.Reply, err error) *RequestsExecutorMock {
	if mmSave.mock.funcSave != nil {
		mmSave.mock.t.Fatalf("RequestsExecutorMock.Save mock is already set by Set")
	}

	if mmSave.defaultExpectation == nil {
		mmSave.defaultExpectation = &RequestsExecutorMockSaveExpectation{mock: mmSave.mock}
	}
	mmSave.defaultExpectation.results = &RequestsExecutorMockSaveResults{r1, err}
	return mmSave.mock
}

//Set uses given function f to mock the RequestsExecutor.Save method
func (mmSave *mRequestsExecutorMockSave) Set(f func(ctx context.Context, current *Transcript, res artifacts.RequestResult) (r1 insolar.Reply, err error)) *RequestsExecutorMock {
	if mmSave.defaultExpectation != nil {
		mmSave.mock.t.Fatalf("Default expectation is already set for the RequestsExecutor.Save method")
	}

	if len(mmSave.expectations) > 0 {
		mmSave.mock.t.Fatalf("Some expectations are already set for the RequestsExecutor.Save method")
	}

	mmSave.mock.funcSave = f
	return mmSave.mock
}

// When sets expectation for the RequestsExecutor.Save which will trigger the result defined by the following
// Then helper
func (mmSave *mRequestsExecutorMockSave) When(ctx context.Context, current *Transcript, res artifacts.RequestResult) *RequestsExecutorMockSaveExpectation {
	if mmSave.mock.funcSave != nil {
		mmSave.mock.t.Fatalf("RequestsExecutorMock.Save mock is already set by Set")
	}

	expectation := &RequestsExecutorMockSaveExpectation{
		mock:   mmSave.mock,
		params: &RequestsExecutorMockSaveParams{ctx, current, res},
	}
	mmSave.expectations = append(mmSave.expectations, expectation)
	return expectation
}

// Then sets up RequestsExecutor.Save return parameters for the expectation previously defined by the When method
func (e *RequestsExecutorMockSaveExpectation) Then(r1 insolar.Reply, err error) *RequestsExecutorMock {
	e.results = &RequestsExecutorMockSaveResults{r1, err}
	return e.mock
}

// Save implements RequestsExecutor
func (mmSave *RequestsExecutorMock) Save(ctx context.Context, current *Transcript, res artifacts.RequestResult) (r1 insolar.Reply, err error) {
	mm_atomic.AddUint64(&mmSave.beforeSaveCounter, 1)
	defer mm_atomic.AddUint64(&mmSave.afterSaveCounter, 1)

	if mmSave.inspectFuncSave != nil {
		mmSave.inspectFuncSave(ctx, current, res)
	}

	params := &RequestsExecutorMockSaveParams{ctx, current, res}

	// Record call args
	mmSave.SaveMock.mutex.Lock()
	mmSave.SaveMock.callArgs = append(mmSave.SaveMock.callArgs, params)
	mmSave.SaveMock.mutex.Unlock()

	for _, e := range mmSave.SaveMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.r1, e.results.err
		}
	}

	if mmSave.SaveMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSave.SaveMock.defaultExpectation.Counter, 1)
		want := mmSave.SaveMock.defaultExpectation.params
		got := RequestsExecutorMockSaveParams{ctx, current, res}
		if want != nil && !minimock.Equal(*want, got) {
			mmSave.t.Errorf("RequestsExecutorMock.Save got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmSave.SaveMock.defaultExpectation.results
		if results == nil {
			mmSave.t.Fatal("No results are set for the RequestsExecutorMock.Save")
		}
		return (*results).r1, (*results).err
	}
	if mmSave.funcSave != nil {
		return mmSave.funcSave(ctx, current, res)
	}
	mmSave.t.Fatalf("Unexpected call to RequestsExecutorMock.Save. %v %v %v", ctx, current, res)
	return
}

// SaveAfterCounter returns a count of finished RequestsExecutorMock.Save invocations
func (mmSave *RequestsExecutorMock) SaveAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSave.afterSaveCounter)
}

// SaveBeforeCounter returns a count of RequestsExecutorMock.Save invocations
func (mmSave *RequestsExecutorMock) SaveBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSave.beforeSaveCounter)
}

// Calls returns a list of arguments used in each call to RequestsExecutorMock.Save.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSave *mRequestsExecutorMockSave) Calls() []*RequestsExecutorMockSaveParams {
	mmSave.mutex.RLock()

	argCopy := make([]*RequestsExecutorMockSaveParams, len(mmSave.callArgs))
	copy(argCopy, mmSave.callArgs)

	mmSave.mutex.RUnlock()

	return argCopy
}

// MinimockSaveDone returns true if the count of the Save invocations corresponds
// the number of defined expectations
func (m *RequestsExecutorMock) MinimockSaveDone() bool {
	for _, e := range m.SaveMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SaveMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSaveCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSave != nil && mm_atomic.LoadUint64(&m.afterSaveCounter) < 1 {
		return false
	}
	return true
}

// MinimockSaveInspect logs each unmet expectation
func (m *RequestsExecutorMock) MinimockSaveInspect() {
	for _, e := range m.SaveMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RequestsExecutorMock.Save with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SaveMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSaveCounter) < 1 {
		if m.SaveMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RequestsExecutorMock.Save")
		} else {
			m.t.Errorf("Expected call to RequestsExecutorMock.Save with params: %#v", *m.SaveMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSave != nil && mm_atomic.LoadUint64(&m.afterSaveCounter) < 1 {
		m.t.Error("Expected call to RequestsExecutorMock.Save")
	}
}

type mRequestsExecutorMockSendReply struct {
	mock               *RequestsExecutorMock
	defaultExpectation *RequestsExecutorMockSendReplyExpectation
	expectations       []*RequestsExecutorMockSendReplyExpectation

	callArgs []*RequestsExecutorMockSendReplyParams
	mutex    sync.RWMutex
}

// RequestsExecutorMockSendReplyExpectation specifies expectation struct of the RequestsExecutor.SendReply
type RequestsExecutorMockSendReplyExpectation struct {
	mock   *RequestsExecutorMock
	params *RequestsExecutorMockSendReplyParams

	Counter uint64
}

// RequestsExecutorMockSendReplyParams contains parameters of the RequestsExecutor.SendReply
type RequestsExecutorMockSendReplyParams struct {
	ctx     context.Context
	current *Transcript
	re      insolar.Reply
	err     error
}

// Expect sets up expected params for RequestsExecutor.SendReply
func (mmSendReply *mRequestsExecutorMockSendReply) Expect(ctx context.Context, current *Transcript, re insolar.Reply, err error) *mRequestsExecutorMockSendReply {
	if mmSendReply.mock.funcSendReply != nil {
		mmSendReply.mock.t.Fatalf("RequestsExecutorMock.SendReply mock is already set by Set")
	}

	if mmSendReply.defaultExpectation == nil {
		mmSendReply.defaultExpectation = &RequestsExecutorMockSendReplyExpectation{}
	}

	mmSendReply.defaultExpectation.params = &RequestsExecutorMockSendReplyParams{ctx, current, re, err}
	for _, e := range mmSendReply.expectations {
		if minimock.Equal(e.params, mmSendReply.defaultExpectation.params) {
			mmSendReply.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSendReply.defaultExpectation.params)
		}
	}

	return mmSendReply
}

// Inspect accepts an inspector function that has same arguments as the RequestsExecutor.SendReply
func (mmSendReply *mRequestsExecutorMockSendReply) Inspect(f func(ctx context.Context, current *Transcript, re insolar.Reply, err error)) *mRequestsExecutorMockSendReply {
	if mmSendReply.mock.inspectFuncSendReply != nil {
		mmSendReply.mock.t.Fatalf("Inspect function is already set for RequestsExecutorMock.SendReply")
	}

	mmSendReply.mock.inspectFuncSendReply = f

	return mmSendReply
}

// Return sets up results that will be returned by RequestsExecutor.SendReply
func (mmSendReply *mRequestsExecutorMockSendReply) Return() *RequestsExecutorMock {
	if mmSendReply.mock.funcSendReply != nil {
		mmSendReply.mock.t.Fatalf("RequestsExecutorMock.SendReply mock is already set by Set")
	}

	if mmSendReply.defaultExpectation == nil {
		mmSendReply.defaultExpectation = &RequestsExecutorMockSendReplyExpectation{mock: mmSendReply.mock}
	}

	return mmSendReply.mock
}

//Set uses given function f to mock the RequestsExecutor.SendReply method
func (mmSendReply *mRequestsExecutorMockSendReply) Set(f func(ctx context.Context, current *Transcript, re insolar.Reply, err error)) *RequestsExecutorMock {
	if mmSendReply.defaultExpectation != nil {
		mmSendReply.mock.t.Fatalf("Default expectation is already set for the RequestsExecutor.SendReply method")
	}

	if len(mmSendReply.expectations) > 0 {
		mmSendReply.mock.t.Fatalf("Some expectations are already set for the RequestsExecutor.SendReply method")
	}

	mmSendReply.mock.funcSendReply = f
	return mmSendReply.mock
}

// SendReply implements RequestsExecutor
func (mmSendReply *RequestsExecutorMock) SendReply(ctx context.Context, current *Transcript, re insolar.Reply, err error) {
	mm_atomic.AddUint64(&mmSendReply.beforeSendReplyCounter, 1)
	defer mm_atomic.AddUint64(&mmSendReply.afterSendReplyCounter, 1)

	if mmSendReply.inspectFuncSendReply != nil {
		mmSendReply.inspectFuncSendReply(ctx, current, re, err)
	}

	params := &RequestsExecutorMockSendReplyParams{ctx, current, re, err}

	// Record call args
	mmSendReply.SendReplyMock.mutex.Lock()
	mmSendReply.SendReplyMock.callArgs = append(mmSendReply.SendReplyMock.callArgs, params)
	mmSendReply.SendReplyMock.mutex.Unlock()

	for _, e := range mmSendReply.SendReplyMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmSendReply.SendReplyMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSendReply.SendReplyMock.defaultExpectation.Counter, 1)
		want := mmSendReply.SendReplyMock.defaultExpectation.params
		got := RequestsExecutorMockSendReplyParams{ctx, current, re, err}
		if want != nil && !minimock.Equal(*want, got) {
			mmSendReply.t.Errorf("RequestsExecutorMock.SendReply got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		return

	}
	if mmSendReply.funcSendReply != nil {
		mmSendReply.funcSendReply(ctx, current, re, err)
		return
	}
	mmSendReply.t.Fatalf("Unexpected call to RequestsExecutorMock.SendReply. %v %v %v %v", ctx, current, re, err)

}

// SendReplyAfterCounter returns a count of finished RequestsExecutorMock.SendReply invocations
func (mmSendReply *RequestsExecutorMock) SendReplyAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendReply.afterSendReplyCounter)
}

// SendReplyBeforeCounter returns a count of RequestsExecutorMock.SendReply invocations
func (mmSendReply *RequestsExecutorMock) SendReplyBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendReply.beforeSendReplyCounter)
}

// Calls returns a list of arguments used in each call to RequestsExecutorMock.SendReply.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSendReply *mRequestsExecutorMockSendReply) Calls() []*RequestsExecutorMockSendReplyParams {
	mmSendReply.mutex.RLock()

	argCopy := make([]*RequestsExecutorMockSendReplyParams, len(mmSendReply.callArgs))
	copy(argCopy, mmSendReply.callArgs)

	mmSendReply.mutex.RUnlock()

	return argCopy
}

// MinimockSendReplyDone returns true if the count of the SendReply invocations corresponds
// the number of defined expectations
func (m *RequestsExecutorMock) MinimockSendReplyDone() bool {
	for _, e := range m.SendReplyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendReplyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendReplyCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendReply != nil && mm_atomic.LoadUint64(&m.afterSendReplyCounter) < 1 {
		return false
	}
	return true
}

// MinimockSendReplyInspect logs each unmet expectation
func (m *RequestsExecutorMock) MinimockSendReplyInspect() {
	for _, e := range m.SendReplyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RequestsExecutorMock.SendReply with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendReplyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendReplyCounter) < 1 {
		if m.SendReplyMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RequestsExecutorMock.SendReply")
		} else {
			m.t.Errorf("Expected call to RequestsExecutorMock.SendReply with params: %#v", *m.SendReplyMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendReply != nil && mm_atomic.LoadUint64(&m.afterSendReplyCounter) < 1 {
		m.t.Error("Expected call to RequestsExecutorMock.SendReply")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *RequestsExecutorMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockExecuteInspect()

		m.MinimockExecuteAndSaveInspect()

		m.MinimockSaveInspect()

		m.MinimockSendReplyInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *RequestsExecutorMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *RequestsExecutorMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockExecuteDone() &&
		m.MinimockExecuteAndSaveDone() &&
		m.MinimockSaveDone() &&
		m.MinimockSendReplyDone()
}
