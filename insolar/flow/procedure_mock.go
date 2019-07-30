package flow

// Code generated by http://github.com/gojuno/minimock (3.0.0). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// ProcedureMock implements Procedure
type ProcedureMock struct {
	t minimock.Tester

	funcProceed          func(ctx context.Context) (err error)
	inspectFuncProceed   func(ctx context.Context)
	afterProceedCounter  uint64
	beforeProceedCounter uint64
	ProceedMock          mProcedureMockProceed
}

// NewProcedureMock returns a mock for Procedure
func NewProcedureMock(t minimock.Tester) *ProcedureMock {
	m := &ProcedureMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ProceedMock = mProcedureMockProceed{mock: m}
	m.ProceedMock.callArgs = []*ProcedureMockProceedParams{}

	return m
}

type mProcedureMockProceed struct {
	mock               *ProcedureMock
	defaultExpectation *ProcedureMockProceedExpectation
	expectations       []*ProcedureMockProceedExpectation

	callArgs []*ProcedureMockProceedParams
	mutex    sync.RWMutex
}

// ProcedureMockProceedExpectation specifies expectation struct of the Procedure.Proceed
type ProcedureMockProceedExpectation struct {
	mock    *ProcedureMock
	params  *ProcedureMockProceedParams
	results *ProcedureMockProceedResults
	Counter uint64
}

// ProcedureMockProceedParams contains parameters of the Procedure.Proceed
type ProcedureMockProceedParams struct {
	ctx context.Context
}

// ProcedureMockProceedResults contains results of the Procedure.Proceed
type ProcedureMockProceedResults struct {
	err error
}

// Expect sets up expected params for Procedure.Proceed
func (mmProceed *mProcedureMockProceed) Expect(ctx context.Context) *mProcedureMockProceed {
	if mmProceed.mock.funcProceed != nil {
		mmProceed.mock.t.Fatalf("ProcedureMock.Proceed mock is already set by Set")
	}

	if mmProceed.defaultExpectation == nil {
		mmProceed.defaultExpectation = &ProcedureMockProceedExpectation{}
	}

	mmProceed.defaultExpectation.params = &ProcedureMockProceedParams{ctx}
	for _, e := range mmProceed.expectations {
		if minimock.Equal(e.params, mmProceed.defaultExpectation.params) {
			mmProceed.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmProceed.defaultExpectation.params)
		}
	}

	return mmProceed
}

// Inspect accepts an inspector function that has same arguments as the Procedure.Proceed
func (mmProceed *mProcedureMockProceed) Inspect(f func(ctx context.Context)) *mProcedureMockProceed {
	if mmProceed.mock.inspectFuncProceed != nil {
		mmProceed.mock.t.Fatalf("Inspect function is already set for ProcedureMock.Proceed")
	}

	mmProceed.mock.inspectFuncProceed = f

	return mmProceed
}

// Return sets up results that will be returned by Procedure.Proceed
func (mmProceed *mProcedureMockProceed) Return(err error) *ProcedureMock {
	if mmProceed.mock.funcProceed != nil {
		mmProceed.mock.t.Fatalf("ProcedureMock.Proceed mock is already set by Set")
	}

	if mmProceed.defaultExpectation == nil {
		mmProceed.defaultExpectation = &ProcedureMockProceedExpectation{mock: mmProceed.mock}
	}
	mmProceed.defaultExpectation.results = &ProcedureMockProceedResults{err}
	return mmProceed.mock
}

//Set uses given function f to mock the Procedure.Proceed method
func (mmProceed *mProcedureMockProceed) Set(f func(ctx context.Context) (err error)) *ProcedureMock {
	if mmProceed.defaultExpectation != nil {
		mmProceed.mock.t.Fatalf("Default expectation is already set for the Procedure.Proceed method")
	}

	if len(mmProceed.expectations) > 0 {
		mmProceed.mock.t.Fatalf("Some expectations are already set for the Procedure.Proceed method")
	}

	mmProceed.mock.funcProceed = f
	return mmProceed.mock
}

// When sets expectation for the Procedure.Proceed which will trigger the result defined by the following
// Then helper
func (mmProceed *mProcedureMockProceed) When(ctx context.Context) *ProcedureMockProceedExpectation {
	if mmProceed.mock.funcProceed != nil {
		mmProceed.mock.t.Fatalf("ProcedureMock.Proceed mock is already set by Set")
	}

	expectation := &ProcedureMockProceedExpectation{
		mock:   mmProceed.mock,
		params: &ProcedureMockProceedParams{ctx},
	}
	mmProceed.expectations = append(mmProceed.expectations, expectation)
	return expectation
}

// Then sets up Procedure.Proceed return parameters for the expectation previously defined by the When method
func (e *ProcedureMockProceedExpectation) Then(err error) *ProcedureMock {
	e.results = &ProcedureMockProceedResults{err}
	return e.mock
}

// Proceed implements Procedure
func (mmProceed *ProcedureMock) Proceed(ctx context.Context) (err error) {
	mm_atomic.AddUint64(&mmProceed.beforeProceedCounter, 1)
	defer mm_atomic.AddUint64(&mmProceed.afterProceedCounter, 1)

	if mmProceed.inspectFuncProceed != nil {
		mmProceed.inspectFuncProceed(ctx)
	}

	params := &ProcedureMockProceedParams{ctx}

	// Record call args
	mmProceed.ProceedMock.mutex.Lock()
	mmProceed.ProceedMock.callArgs = append(mmProceed.ProceedMock.callArgs, params)
	mmProceed.ProceedMock.mutex.Unlock()

	for _, e := range mmProceed.ProceedMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmProceed.ProceedMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmProceed.ProceedMock.defaultExpectation.Counter, 1)
		want := mmProceed.ProceedMock.defaultExpectation.params
		got := ProcedureMockProceedParams{ctx}
		if want != nil && !minimock.Equal(*want, got) {
			mmProceed.t.Errorf("ProcedureMock.Proceed got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmProceed.ProceedMock.defaultExpectation.results
		if results == nil {
			mmProceed.t.Fatal("No results are set for the ProcedureMock.Proceed")
		}
		return (*results).err
	}
	if mmProceed.funcProceed != nil {
		return mmProceed.funcProceed(ctx)
	}
	mmProceed.t.Fatalf("Unexpected call to ProcedureMock.Proceed. %v", ctx)
	return
}

// ProceedAfterCounter returns a count of finished ProcedureMock.Proceed invocations
func (mmProceed *ProcedureMock) ProceedAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmProceed.afterProceedCounter)
}

// ProceedBeforeCounter returns a count of ProcedureMock.Proceed invocations
func (mmProceed *ProcedureMock) ProceedBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmProceed.beforeProceedCounter)
}

// Calls returns a list of arguments used in each call to ProcedureMock.Proceed.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmProceed *mProcedureMockProceed) Calls() []*ProcedureMockProceedParams {
	mmProceed.mutex.RLock()

	argCopy := make([]*ProcedureMockProceedParams, len(mmProceed.callArgs))
	copy(argCopy, mmProceed.callArgs)

	mmProceed.mutex.RUnlock()

	return argCopy
}

// MinimockProceedDone returns true if the count of the Proceed invocations corresponds
// the number of defined expectations
func (m *ProcedureMock) MinimockProceedDone() bool {
	for _, e := range m.ProceedMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ProceedMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterProceedCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcProceed != nil && mm_atomic.LoadUint64(&m.afterProceedCounter) < 1 {
		return false
	}
	return true
}

// MinimockProceedInspect logs each unmet expectation
func (m *ProcedureMock) MinimockProceedInspect() {
	for _, e := range m.ProceedMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ProcedureMock.Proceed with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ProceedMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterProceedCounter) < 1 {
		if m.ProceedMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ProcedureMock.Proceed")
		} else {
			m.t.Errorf("Expected call to ProcedureMock.Proceed with params: %#v", *m.ProceedMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcProceed != nil && mm_atomic.LoadUint64(&m.afterProceedCounter) < 1 {
		m.t.Error("Expected call to ProcedureMock.Proceed")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ProcedureMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockProceedInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ProcedureMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *ProcedureMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockProceedDone()
}
