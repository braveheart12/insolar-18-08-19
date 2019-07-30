package misbehavior

// Code generated by http://github.com/gojuno/minimock (3.0.0). DO NOT EDIT.

import (
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/insolar/insolar/network/consensus/common/endpoints"
	"github.com/insolar/insolar/network/consensus/gcpv2/api/profiles"
)

// ReportMock implements Report
type ReportMock struct {
	t minimock.Tester

	funcCaptureMark          func() (p1 interface{})
	inspectFuncCaptureMark   func()
	afterCaptureMarkCounter  uint64
	beforeCaptureMarkCounter uint64
	CaptureMarkMock          mReportMockCaptureMark

	funcDetails          func() (pa1 []interface{})
	inspectFuncDetails   func()
	afterDetailsCounter  uint64
	beforeDetailsCounter uint64
	DetailsMock          mReportMockDetails

	funcMisbehaviorType          func() (t1 Type)
	inspectFuncMisbehaviorType   func()
	afterMisbehaviorTypeCounter  uint64
	beforeMisbehaviorTypeCounter uint64
	MisbehaviorTypeMock          mReportMockMisbehaviorType

	funcViolatorHost          func() (i1 endpoints.InboundConnection)
	inspectFuncViolatorHost   func()
	afterViolatorHostCounter  uint64
	beforeViolatorHostCounter uint64
	ViolatorHostMock          mReportMockViolatorHost

	funcViolatorNode          func() (b1 profiles.BaseNode)
	inspectFuncViolatorNode   func()
	afterViolatorNodeCounter  uint64
	beforeViolatorNodeCounter uint64
	ViolatorNodeMock          mReportMockViolatorNode
}

// NewReportMock returns a mock for Report
func NewReportMock(t minimock.Tester) *ReportMock {
	m := &ReportMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CaptureMarkMock = mReportMockCaptureMark{mock: m}

	m.DetailsMock = mReportMockDetails{mock: m}

	m.MisbehaviorTypeMock = mReportMockMisbehaviorType{mock: m}

	m.ViolatorHostMock = mReportMockViolatorHost{mock: m}

	m.ViolatorNodeMock = mReportMockViolatorNode{mock: m}

	return m
}

type mReportMockCaptureMark struct {
	mock               *ReportMock
	defaultExpectation *ReportMockCaptureMarkExpectation
	expectations       []*ReportMockCaptureMarkExpectation
}

// ReportMockCaptureMarkExpectation specifies expectation struct of the Report.CaptureMark
type ReportMockCaptureMarkExpectation struct {
	mock *ReportMock

	results *ReportMockCaptureMarkResults
	Counter uint64
}

// ReportMockCaptureMarkResults contains results of the Report.CaptureMark
type ReportMockCaptureMarkResults struct {
	p1 interface{}
}

// Expect sets up expected params for Report.CaptureMark
func (mmCaptureMark *mReportMockCaptureMark) Expect() *mReportMockCaptureMark {
	if mmCaptureMark.mock.funcCaptureMark != nil {
		mmCaptureMark.mock.t.Fatalf("ReportMock.CaptureMark mock is already set by Set")
	}

	if mmCaptureMark.defaultExpectation == nil {
		mmCaptureMark.defaultExpectation = &ReportMockCaptureMarkExpectation{}
	}

	return mmCaptureMark
}

// Inspect accepts an inspector function that has same arguments as the Report.CaptureMark
func (mmCaptureMark *mReportMockCaptureMark) Inspect(f func()) *mReportMockCaptureMark {
	if mmCaptureMark.mock.inspectFuncCaptureMark != nil {
		mmCaptureMark.mock.t.Fatalf("Inspect function is already set for ReportMock.CaptureMark")
	}

	mmCaptureMark.mock.inspectFuncCaptureMark = f

	return mmCaptureMark
}

// Return sets up results that will be returned by Report.CaptureMark
func (mmCaptureMark *mReportMockCaptureMark) Return(p1 interface{}) *ReportMock {
	if mmCaptureMark.mock.funcCaptureMark != nil {
		mmCaptureMark.mock.t.Fatalf("ReportMock.CaptureMark mock is already set by Set")
	}

	if mmCaptureMark.defaultExpectation == nil {
		mmCaptureMark.defaultExpectation = &ReportMockCaptureMarkExpectation{mock: mmCaptureMark.mock}
	}
	mmCaptureMark.defaultExpectation.results = &ReportMockCaptureMarkResults{p1}
	return mmCaptureMark.mock
}

//Set uses given function f to mock the Report.CaptureMark method
func (mmCaptureMark *mReportMockCaptureMark) Set(f func() (p1 interface{})) *ReportMock {
	if mmCaptureMark.defaultExpectation != nil {
		mmCaptureMark.mock.t.Fatalf("Default expectation is already set for the Report.CaptureMark method")
	}

	if len(mmCaptureMark.expectations) > 0 {
		mmCaptureMark.mock.t.Fatalf("Some expectations are already set for the Report.CaptureMark method")
	}

	mmCaptureMark.mock.funcCaptureMark = f
	return mmCaptureMark.mock
}

// CaptureMark implements Report
func (mmCaptureMark *ReportMock) CaptureMark() (p1 interface{}) {
	mm_atomic.AddUint64(&mmCaptureMark.beforeCaptureMarkCounter, 1)
	defer mm_atomic.AddUint64(&mmCaptureMark.afterCaptureMarkCounter, 1)

	if mmCaptureMark.inspectFuncCaptureMark != nil {
		mmCaptureMark.inspectFuncCaptureMark()
	}

	if mmCaptureMark.CaptureMarkMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCaptureMark.CaptureMarkMock.defaultExpectation.Counter, 1)

		results := mmCaptureMark.CaptureMarkMock.defaultExpectation.results
		if results == nil {
			mmCaptureMark.t.Fatal("No results are set for the ReportMock.CaptureMark")
		}
		return (*results).p1
	}
	if mmCaptureMark.funcCaptureMark != nil {
		return mmCaptureMark.funcCaptureMark()
	}
	mmCaptureMark.t.Fatalf("Unexpected call to ReportMock.CaptureMark.")
	return
}

// CaptureMarkAfterCounter returns a count of finished ReportMock.CaptureMark invocations
func (mmCaptureMark *ReportMock) CaptureMarkAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCaptureMark.afterCaptureMarkCounter)
}

// CaptureMarkBeforeCounter returns a count of ReportMock.CaptureMark invocations
func (mmCaptureMark *ReportMock) CaptureMarkBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCaptureMark.beforeCaptureMarkCounter)
}

// MinimockCaptureMarkDone returns true if the count of the CaptureMark invocations corresponds
// the number of defined expectations
func (m *ReportMock) MinimockCaptureMarkDone() bool {
	for _, e := range m.CaptureMarkMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CaptureMarkMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCaptureMarkCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCaptureMark != nil && mm_atomic.LoadUint64(&m.afterCaptureMarkCounter) < 1 {
		return false
	}
	return true
}

// MinimockCaptureMarkInspect logs each unmet expectation
func (m *ReportMock) MinimockCaptureMarkInspect() {
	for _, e := range m.CaptureMarkMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to ReportMock.CaptureMark")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CaptureMarkMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCaptureMarkCounter) < 1 {
		m.t.Error("Expected call to ReportMock.CaptureMark")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCaptureMark != nil && mm_atomic.LoadUint64(&m.afterCaptureMarkCounter) < 1 {
		m.t.Error("Expected call to ReportMock.CaptureMark")
	}
}

type mReportMockDetails struct {
	mock               *ReportMock
	defaultExpectation *ReportMockDetailsExpectation
	expectations       []*ReportMockDetailsExpectation
}

// ReportMockDetailsExpectation specifies expectation struct of the Report.Details
type ReportMockDetailsExpectation struct {
	mock *ReportMock

	results *ReportMockDetailsResults
	Counter uint64
}

// ReportMockDetailsResults contains results of the Report.Details
type ReportMockDetailsResults struct {
	pa1 []interface{}
}

// Expect sets up expected params for Report.Details
func (mmDetails *mReportMockDetails) Expect() *mReportMockDetails {
	if mmDetails.mock.funcDetails != nil {
		mmDetails.mock.t.Fatalf("ReportMock.Details mock is already set by Set")
	}

	if mmDetails.defaultExpectation == nil {
		mmDetails.defaultExpectation = &ReportMockDetailsExpectation{}
	}

	return mmDetails
}

// Inspect accepts an inspector function that has same arguments as the Report.Details
func (mmDetails *mReportMockDetails) Inspect(f func()) *mReportMockDetails {
	if mmDetails.mock.inspectFuncDetails != nil {
		mmDetails.mock.t.Fatalf("Inspect function is already set for ReportMock.Details")
	}

	mmDetails.mock.inspectFuncDetails = f

	return mmDetails
}

// Return sets up results that will be returned by Report.Details
func (mmDetails *mReportMockDetails) Return(pa1 []interface{}) *ReportMock {
	if mmDetails.mock.funcDetails != nil {
		mmDetails.mock.t.Fatalf("ReportMock.Details mock is already set by Set")
	}

	if mmDetails.defaultExpectation == nil {
		mmDetails.defaultExpectation = &ReportMockDetailsExpectation{mock: mmDetails.mock}
	}
	mmDetails.defaultExpectation.results = &ReportMockDetailsResults{pa1}
	return mmDetails.mock
}

//Set uses given function f to mock the Report.Details method
func (mmDetails *mReportMockDetails) Set(f func() (pa1 []interface{})) *ReportMock {
	if mmDetails.defaultExpectation != nil {
		mmDetails.mock.t.Fatalf("Default expectation is already set for the Report.Details method")
	}

	if len(mmDetails.expectations) > 0 {
		mmDetails.mock.t.Fatalf("Some expectations are already set for the Report.Details method")
	}

	mmDetails.mock.funcDetails = f
	return mmDetails.mock
}

// Details implements Report
func (mmDetails *ReportMock) Details() (pa1 []interface{}) {
	mm_atomic.AddUint64(&mmDetails.beforeDetailsCounter, 1)
	defer mm_atomic.AddUint64(&mmDetails.afterDetailsCounter, 1)

	if mmDetails.inspectFuncDetails != nil {
		mmDetails.inspectFuncDetails()
	}

	if mmDetails.DetailsMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDetails.DetailsMock.defaultExpectation.Counter, 1)

		results := mmDetails.DetailsMock.defaultExpectation.results
		if results == nil {
			mmDetails.t.Fatal("No results are set for the ReportMock.Details")
		}
		return (*results).pa1
	}
	if mmDetails.funcDetails != nil {
		return mmDetails.funcDetails()
	}
	mmDetails.t.Fatalf("Unexpected call to ReportMock.Details.")
	return
}

// DetailsAfterCounter returns a count of finished ReportMock.Details invocations
func (mmDetails *ReportMock) DetailsAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDetails.afterDetailsCounter)
}

// DetailsBeforeCounter returns a count of ReportMock.Details invocations
func (mmDetails *ReportMock) DetailsBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDetails.beforeDetailsCounter)
}

// MinimockDetailsDone returns true if the count of the Details invocations corresponds
// the number of defined expectations
func (m *ReportMock) MinimockDetailsDone() bool {
	for _, e := range m.DetailsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DetailsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDetailsCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDetails != nil && mm_atomic.LoadUint64(&m.afterDetailsCounter) < 1 {
		return false
	}
	return true
}

// MinimockDetailsInspect logs each unmet expectation
func (m *ReportMock) MinimockDetailsInspect() {
	for _, e := range m.DetailsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to ReportMock.Details")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DetailsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDetailsCounter) < 1 {
		m.t.Error("Expected call to ReportMock.Details")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDetails != nil && mm_atomic.LoadUint64(&m.afterDetailsCounter) < 1 {
		m.t.Error("Expected call to ReportMock.Details")
	}
}

type mReportMockMisbehaviorType struct {
	mock               *ReportMock
	defaultExpectation *ReportMockMisbehaviorTypeExpectation
	expectations       []*ReportMockMisbehaviorTypeExpectation
}

// ReportMockMisbehaviorTypeExpectation specifies expectation struct of the Report.MisbehaviorType
type ReportMockMisbehaviorTypeExpectation struct {
	mock *ReportMock

	results *ReportMockMisbehaviorTypeResults
	Counter uint64
}

// ReportMockMisbehaviorTypeResults contains results of the Report.MisbehaviorType
type ReportMockMisbehaviorTypeResults struct {
	t1 Type
}

// Expect sets up expected params for Report.MisbehaviorType
func (mmMisbehaviorType *mReportMockMisbehaviorType) Expect() *mReportMockMisbehaviorType {
	if mmMisbehaviorType.mock.funcMisbehaviorType != nil {
		mmMisbehaviorType.mock.t.Fatalf("ReportMock.MisbehaviorType mock is already set by Set")
	}

	if mmMisbehaviorType.defaultExpectation == nil {
		mmMisbehaviorType.defaultExpectation = &ReportMockMisbehaviorTypeExpectation{}
	}

	return mmMisbehaviorType
}

// Inspect accepts an inspector function that has same arguments as the Report.MisbehaviorType
func (mmMisbehaviorType *mReportMockMisbehaviorType) Inspect(f func()) *mReportMockMisbehaviorType {
	if mmMisbehaviorType.mock.inspectFuncMisbehaviorType != nil {
		mmMisbehaviorType.mock.t.Fatalf("Inspect function is already set for ReportMock.MisbehaviorType")
	}

	mmMisbehaviorType.mock.inspectFuncMisbehaviorType = f

	return mmMisbehaviorType
}

// Return sets up results that will be returned by Report.MisbehaviorType
func (mmMisbehaviorType *mReportMockMisbehaviorType) Return(t1 Type) *ReportMock {
	if mmMisbehaviorType.mock.funcMisbehaviorType != nil {
		mmMisbehaviorType.mock.t.Fatalf("ReportMock.MisbehaviorType mock is already set by Set")
	}

	if mmMisbehaviorType.defaultExpectation == nil {
		mmMisbehaviorType.defaultExpectation = &ReportMockMisbehaviorTypeExpectation{mock: mmMisbehaviorType.mock}
	}
	mmMisbehaviorType.defaultExpectation.results = &ReportMockMisbehaviorTypeResults{t1}
	return mmMisbehaviorType.mock
}

//Set uses given function f to mock the Report.MisbehaviorType method
func (mmMisbehaviorType *mReportMockMisbehaviorType) Set(f func() (t1 Type)) *ReportMock {
	if mmMisbehaviorType.defaultExpectation != nil {
		mmMisbehaviorType.mock.t.Fatalf("Default expectation is already set for the Report.MisbehaviorType method")
	}

	if len(mmMisbehaviorType.expectations) > 0 {
		mmMisbehaviorType.mock.t.Fatalf("Some expectations are already set for the Report.MisbehaviorType method")
	}

	mmMisbehaviorType.mock.funcMisbehaviorType = f
	return mmMisbehaviorType.mock
}

// MisbehaviorType implements Report
func (mmMisbehaviorType *ReportMock) MisbehaviorType() (t1 Type) {
	mm_atomic.AddUint64(&mmMisbehaviorType.beforeMisbehaviorTypeCounter, 1)
	defer mm_atomic.AddUint64(&mmMisbehaviorType.afterMisbehaviorTypeCounter, 1)

	if mmMisbehaviorType.inspectFuncMisbehaviorType != nil {
		mmMisbehaviorType.inspectFuncMisbehaviorType()
	}

	if mmMisbehaviorType.MisbehaviorTypeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmMisbehaviorType.MisbehaviorTypeMock.defaultExpectation.Counter, 1)

		results := mmMisbehaviorType.MisbehaviorTypeMock.defaultExpectation.results
		if results == nil {
			mmMisbehaviorType.t.Fatal("No results are set for the ReportMock.MisbehaviorType")
		}
		return (*results).t1
	}
	if mmMisbehaviorType.funcMisbehaviorType != nil {
		return mmMisbehaviorType.funcMisbehaviorType()
	}
	mmMisbehaviorType.t.Fatalf("Unexpected call to ReportMock.MisbehaviorType.")
	return
}

// MisbehaviorTypeAfterCounter returns a count of finished ReportMock.MisbehaviorType invocations
func (mmMisbehaviorType *ReportMock) MisbehaviorTypeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmMisbehaviorType.afterMisbehaviorTypeCounter)
}

// MisbehaviorTypeBeforeCounter returns a count of ReportMock.MisbehaviorType invocations
func (mmMisbehaviorType *ReportMock) MisbehaviorTypeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmMisbehaviorType.beforeMisbehaviorTypeCounter)
}

// MinimockMisbehaviorTypeDone returns true if the count of the MisbehaviorType invocations corresponds
// the number of defined expectations
func (m *ReportMock) MinimockMisbehaviorTypeDone() bool {
	for _, e := range m.MisbehaviorTypeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.MisbehaviorTypeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterMisbehaviorTypeCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcMisbehaviorType != nil && mm_atomic.LoadUint64(&m.afterMisbehaviorTypeCounter) < 1 {
		return false
	}
	return true
}

// MinimockMisbehaviorTypeInspect logs each unmet expectation
func (m *ReportMock) MinimockMisbehaviorTypeInspect() {
	for _, e := range m.MisbehaviorTypeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to ReportMock.MisbehaviorType")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.MisbehaviorTypeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterMisbehaviorTypeCounter) < 1 {
		m.t.Error("Expected call to ReportMock.MisbehaviorType")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcMisbehaviorType != nil && mm_atomic.LoadUint64(&m.afterMisbehaviorTypeCounter) < 1 {
		m.t.Error("Expected call to ReportMock.MisbehaviorType")
	}
}

type mReportMockViolatorHost struct {
	mock               *ReportMock
	defaultExpectation *ReportMockViolatorHostExpectation
	expectations       []*ReportMockViolatorHostExpectation
}

// ReportMockViolatorHostExpectation specifies expectation struct of the Report.ViolatorHost
type ReportMockViolatorHostExpectation struct {
	mock *ReportMock

	results *ReportMockViolatorHostResults
	Counter uint64
}

// ReportMockViolatorHostResults contains results of the Report.ViolatorHost
type ReportMockViolatorHostResults struct {
	i1 endpoints.InboundConnection
}

// Expect sets up expected params for Report.ViolatorHost
func (mmViolatorHost *mReportMockViolatorHost) Expect() *mReportMockViolatorHost {
	if mmViolatorHost.mock.funcViolatorHost != nil {
		mmViolatorHost.mock.t.Fatalf("ReportMock.ViolatorHost mock is already set by Set")
	}

	if mmViolatorHost.defaultExpectation == nil {
		mmViolatorHost.defaultExpectation = &ReportMockViolatorHostExpectation{}
	}

	return mmViolatorHost
}

// Inspect accepts an inspector function that has same arguments as the Report.ViolatorHost
func (mmViolatorHost *mReportMockViolatorHost) Inspect(f func()) *mReportMockViolatorHost {
	if mmViolatorHost.mock.inspectFuncViolatorHost != nil {
		mmViolatorHost.mock.t.Fatalf("Inspect function is already set for ReportMock.ViolatorHost")
	}

	mmViolatorHost.mock.inspectFuncViolatorHost = f

	return mmViolatorHost
}

// Return sets up results that will be returned by Report.ViolatorHost
func (mmViolatorHost *mReportMockViolatorHost) Return(i1 endpoints.InboundConnection) *ReportMock {
	if mmViolatorHost.mock.funcViolatorHost != nil {
		mmViolatorHost.mock.t.Fatalf("ReportMock.ViolatorHost mock is already set by Set")
	}

	if mmViolatorHost.defaultExpectation == nil {
		mmViolatorHost.defaultExpectation = &ReportMockViolatorHostExpectation{mock: mmViolatorHost.mock}
	}
	mmViolatorHost.defaultExpectation.results = &ReportMockViolatorHostResults{i1}
	return mmViolatorHost.mock
}

//Set uses given function f to mock the Report.ViolatorHost method
func (mmViolatorHost *mReportMockViolatorHost) Set(f func() (i1 endpoints.InboundConnection)) *ReportMock {
	if mmViolatorHost.defaultExpectation != nil {
		mmViolatorHost.mock.t.Fatalf("Default expectation is already set for the Report.ViolatorHost method")
	}

	if len(mmViolatorHost.expectations) > 0 {
		mmViolatorHost.mock.t.Fatalf("Some expectations are already set for the Report.ViolatorHost method")
	}

	mmViolatorHost.mock.funcViolatorHost = f
	return mmViolatorHost.mock
}

// ViolatorHost implements Report
func (mmViolatorHost *ReportMock) ViolatorHost() (i1 endpoints.InboundConnection) {
	mm_atomic.AddUint64(&mmViolatorHost.beforeViolatorHostCounter, 1)
	defer mm_atomic.AddUint64(&mmViolatorHost.afterViolatorHostCounter, 1)

	if mmViolatorHost.inspectFuncViolatorHost != nil {
		mmViolatorHost.inspectFuncViolatorHost()
	}

	if mmViolatorHost.ViolatorHostMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmViolatorHost.ViolatorHostMock.defaultExpectation.Counter, 1)

		results := mmViolatorHost.ViolatorHostMock.defaultExpectation.results
		if results == nil {
			mmViolatorHost.t.Fatal("No results are set for the ReportMock.ViolatorHost")
		}
		return (*results).i1
	}
	if mmViolatorHost.funcViolatorHost != nil {
		return mmViolatorHost.funcViolatorHost()
	}
	mmViolatorHost.t.Fatalf("Unexpected call to ReportMock.ViolatorHost.")
	return
}

// ViolatorHostAfterCounter returns a count of finished ReportMock.ViolatorHost invocations
func (mmViolatorHost *ReportMock) ViolatorHostAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmViolatorHost.afterViolatorHostCounter)
}

// ViolatorHostBeforeCounter returns a count of ReportMock.ViolatorHost invocations
func (mmViolatorHost *ReportMock) ViolatorHostBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmViolatorHost.beforeViolatorHostCounter)
}

// MinimockViolatorHostDone returns true if the count of the ViolatorHost invocations corresponds
// the number of defined expectations
func (m *ReportMock) MinimockViolatorHostDone() bool {
	for _, e := range m.ViolatorHostMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ViolatorHostMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterViolatorHostCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcViolatorHost != nil && mm_atomic.LoadUint64(&m.afterViolatorHostCounter) < 1 {
		return false
	}
	return true
}

// MinimockViolatorHostInspect logs each unmet expectation
func (m *ReportMock) MinimockViolatorHostInspect() {
	for _, e := range m.ViolatorHostMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to ReportMock.ViolatorHost")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ViolatorHostMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterViolatorHostCounter) < 1 {
		m.t.Error("Expected call to ReportMock.ViolatorHost")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcViolatorHost != nil && mm_atomic.LoadUint64(&m.afterViolatorHostCounter) < 1 {
		m.t.Error("Expected call to ReportMock.ViolatorHost")
	}
}

type mReportMockViolatorNode struct {
	mock               *ReportMock
	defaultExpectation *ReportMockViolatorNodeExpectation
	expectations       []*ReportMockViolatorNodeExpectation
}

// ReportMockViolatorNodeExpectation specifies expectation struct of the Report.ViolatorNode
type ReportMockViolatorNodeExpectation struct {
	mock *ReportMock

	results *ReportMockViolatorNodeResults
	Counter uint64
}

// ReportMockViolatorNodeResults contains results of the Report.ViolatorNode
type ReportMockViolatorNodeResults struct {
	b1 profiles.BaseNode
}

// Expect sets up expected params for Report.ViolatorNode
func (mmViolatorNode *mReportMockViolatorNode) Expect() *mReportMockViolatorNode {
	if mmViolatorNode.mock.funcViolatorNode != nil {
		mmViolatorNode.mock.t.Fatalf("ReportMock.ViolatorNode mock is already set by Set")
	}

	if mmViolatorNode.defaultExpectation == nil {
		mmViolatorNode.defaultExpectation = &ReportMockViolatorNodeExpectation{}
	}

	return mmViolatorNode
}

// Inspect accepts an inspector function that has same arguments as the Report.ViolatorNode
func (mmViolatorNode *mReportMockViolatorNode) Inspect(f func()) *mReportMockViolatorNode {
	if mmViolatorNode.mock.inspectFuncViolatorNode != nil {
		mmViolatorNode.mock.t.Fatalf("Inspect function is already set for ReportMock.ViolatorNode")
	}

	mmViolatorNode.mock.inspectFuncViolatorNode = f

	return mmViolatorNode
}

// Return sets up results that will be returned by Report.ViolatorNode
func (mmViolatorNode *mReportMockViolatorNode) Return(b1 profiles.BaseNode) *ReportMock {
	if mmViolatorNode.mock.funcViolatorNode != nil {
		mmViolatorNode.mock.t.Fatalf("ReportMock.ViolatorNode mock is already set by Set")
	}

	if mmViolatorNode.defaultExpectation == nil {
		mmViolatorNode.defaultExpectation = &ReportMockViolatorNodeExpectation{mock: mmViolatorNode.mock}
	}
	mmViolatorNode.defaultExpectation.results = &ReportMockViolatorNodeResults{b1}
	return mmViolatorNode.mock
}

//Set uses given function f to mock the Report.ViolatorNode method
func (mmViolatorNode *mReportMockViolatorNode) Set(f func() (b1 profiles.BaseNode)) *ReportMock {
	if mmViolatorNode.defaultExpectation != nil {
		mmViolatorNode.mock.t.Fatalf("Default expectation is already set for the Report.ViolatorNode method")
	}

	if len(mmViolatorNode.expectations) > 0 {
		mmViolatorNode.mock.t.Fatalf("Some expectations are already set for the Report.ViolatorNode method")
	}

	mmViolatorNode.mock.funcViolatorNode = f
	return mmViolatorNode.mock
}

// ViolatorNode implements Report
func (mmViolatorNode *ReportMock) ViolatorNode() (b1 profiles.BaseNode) {
	mm_atomic.AddUint64(&mmViolatorNode.beforeViolatorNodeCounter, 1)
	defer mm_atomic.AddUint64(&mmViolatorNode.afterViolatorNodeCounter, 1)

	if mmViolatorNode.inspectFuncViolatorNode != nil {
		mmViolatorNode.inspectFuncViolatorNode()
	}

	if mmViolatorNode.ViolatorNodeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmViolatorNode.ViolatorNodeMock.defaultExpectation.Counter, 1)

		results := mmViolatorNode.ViolatorNodeMock.defaultExpectation.results
		if results == nil {
			mmViolatorNode.t.Fatal("No results are set for the ReportMock.ViolatorNode")
		}
		return (*results).b1
	}
	if mmViolatorNode.funcViolatorNode != nil {
		return mmViolatorNode.funcViolatorNode()
	}
	mmViolatorNode.t.Fatalf("Unexpected call to ReportMock.ViolatorNode.")
	return
}

// ViolatorNodeAfterCounter returns a count of finished ReportMock.ViolatorNode invocations
func (mmViolatorNode *ReportMock) ViolatorNodeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmViolatorNode.afterViolatorNodeCounter)
}

// ViolatorNodeBeforeCounter returns a count of ReportMock.ViolatorNode invocations
func (mmViolatorNode *ReportMock) ViolatorNodeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmViolatorNode.beforeViolatorNodeCounter)
}

// MinimockViolatorNodeDone returns true if the count of the ViolatorNode invocations corresponds
// the number of defined expectations
func (m *ReportMock) MinimockViolatorNodeDone() bool {
	for _, e := range m.ViolatorNodeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ViolatorNodeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterViolatorNodeCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcViolatorNode != nil && mm_atomic.LoadUint64(&m.afterViolatorNodeCounter) < 1 {
		return false
	}
	return true
}

// MinimockViolatorNodeInspect logs each unmet expectation
func (m *ReportMock) MinimockViolatorNodeInspect() {
	for _, e := range m.ViolatorNodeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to ReportMock.ViolatorNode")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ViolatorNodeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterViolatorNodeCounter) < 1 {
		m.t.Error("Expected call to ReportMock.ViolatorNode")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcViolatorNode != nil && mm_atomic.LoadUint64(&m.afterViolatorNodeCounter) < 1 {
		m.t.Error("Expected call to ReportMock.ViolatorNode")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ReportMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCaptureMarkInspect()

		m.MinimockDetailsInspect()

		m.MinimockMisbehaviorTypeInspect()

		m.MinimockViolatorHostInspect()

		m.MinimockViolatorNodeInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ReportMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *ReportMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCaptureMarkDone() &&
		m.MinimockDetailsDone() &&
		m.MinimockMisbehaviorTypeDone() &&
		m.MinimockViolatorHostDone() &&
		m.MinimockViolatorNodeDone()
}
