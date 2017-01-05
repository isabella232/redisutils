// This file was generated by counterfeiter
// counterfeiter -o monit/fakes/monit.go --fake-name FakeMonit monit/monit.go Monit

package fakes

import (
	"sync"

	"github.com/pivotal-cf/redisutils/monit"
)

//FakeMonit ...
type FakeMonit struct {
	GetSummaryStub        func() (monit.Statuses, error)
	getSummaryMutex       sync.RWMutex
	getSummaryArgsForCall []struct{}
	getSummaryReturns     struct {
		result1 monit.Statuses
		result2 error
	}
	GetStatusStub        func(string) (monit.Status, error)
	getStatusMutex       sync.RWMutex
	getStatusArgsForCall []struct {
		arg1 string
	}
	getStatusReturns struct {
		result1 monit.Status
		result2 error
	}
	StartStub        func(string) error
	startMutex       sync.RWMutex
	startArgsForCall []struct {
		arg1 string
	}
	startReturns struct {
		result1 error
	}
	StopStub        func(string) error
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
		arg1 string
	}
	stopReturns struct {
		result1 error
	}
	StartAndWaitStub        func(string) error
	startAndWaitMutex       sync.RWMutex
	startAndWaitArgsForCall []struct {
		arg1 string
	}
	startAndWaitReturns struct {
		result1 error
	}
	StopAndWaitStub        func(string) error
	stopAndWaitMutex       sync.RWMutex
	stopAndWaitArgsForCall []struct {
		arg1 string
	}
	stopAndWaitReturns struct {
		result1 error
	}
	SetMonitrcPathStub        func(string)
	setMonitrcPathMutex       sync.RWMutex
	setMonitrcPathArgsForCall []struct {
		arg1 string
	}
	SetExecutableStub        func(string)
	setExecutableMutex       sync.RWMutex
	setExecutableArgsForCall []struct {
		arg1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

//GetSummary ...
func (fake *FakeMonit) GetSummary() (monit.Statuses, error) {
	fake.getSummaryMutex.Lock()
	fake.getSummaryArgsForCall = append(fake.getSummaryArgsForCall, struct{}{})
	fake.recordInvocation("GetSummary", []interface{}{})
	fake.getSummaryMutex.Unlock()
	if fake.GetSummaryStub != nil {
		return fake.GetSummaryStub()
	}
	return fake.getSummaryReturns.result1, fake.getSummaryReturns.result2
}

//GetSummaryCallCount ...
func (fake *FakeMonit) GetSummaryCallCount() int {
	fake.getSummaryMutex.RLock()
	defer fake.getSummaryMutex.RUnlock()
	return len(fake.getSummaryArgsForCall)
}

//GetSummaryReturns ...
func (fake *FakeMonit) GetSummaryReturns(result1 monit.Statuses, result2 error) {
	fake.GetSummaryStub = nil
	fake.getSummaryReturns = struct {
		result1 monit.Statuses
		result2 error
	}{result1, result2}
}

//GetStatus ...
func (fake *FakeMonit) GetStatus(arg1 string) (monit.Status, error) {
	fake.getStatusMutex.Lock()
	fake.getStatusArgsForCall = append(fake.getStatusArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetStatus", []interface{}{arg1})
	fake.getStatusMutex.Unlock()
	if fake.GetStatusStub != nil {
		return fake.GetStatusStub(arg1)
	}
	return fake.getStatusReturns.result1, fake.getStatusReturns.result2
}

//GetStatusCallCount ...
func (fake *FakeMonit) GetStatusCallCount() int {
	fake.getStatusMutex.RLock()
	defer fake.getStatusMutex.RUnlock()
	return len(fake.getStatusArgsForCall)
}

//GetStatusArgsForCall ...
func (fake *FakeMonit) GetStatusArgsForCall(i int) string {
	fake.getStatusMutex.RLock()
	defer fake.getStatusMutex.RUnlock()
	return fake.getStatusArgsForCall[i].arg1
}

//GetStatusReturns ...
func (fake *FakeMonit) GetStatusReturns(result1 monit.Status, result2 error) {
	fake.GetStatusStub = nil
	fake.getStatusReturns = struct {
		result1 monit.Status
		result2 error
	}{result1, result2}
}

//Start ...
func (fake *FakeMonit) Start(arg1 string) error {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Start", []interface{}{arg1})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub(arg1)
	}
	return fake.startReturns.result1
}

//StartCallCount ...
func (fake *FakeMonit) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

//StartArgsForCall ...
func (fake *FakeMonit) StartArgsForCall(i int) string {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return fake.startArgsForCall[i].arg1
}

//StartReturns ...
func (fake *FakeMonit) StartReturns(result1 error) {
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

//Stop ...
func (fake *FakeMonit) Stop(arg1 string) error {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Stop", []interface{}{arg1})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		return fake.StopStub(arg1)
	}
	return fake.stopReturns.result1
}

//StopCallCount ...
func (fake *FakeMonit) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

//StopArgsForCall ...
func (fake *FakeMonit) StopArgsForCall(i int) string {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return fake.stopArgsForCall[i].arg1
}

//StopReturns ...
func (fake *FakeMonit) StopReturns(result1 error) {
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

//StartAndWait ...
func (fake *FakeMonit) StartAndWait(arg1 string) error {
	fake.startAndWaitMutex.Lock()
	fake.startAndWaitArgsForCall = append(fake.startAndWaitArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("StartAndWait", []interface{}{arg1})
	fake.startAndWaitMutex.Unlock()
	if fake.StartAndWaitStub != nil {
		return fake.StartAndWaitStub(arg1)
	}
	return fake.startAndWaitReturns.result1
}

//StartAndWaitCallCount ...
func (fake *FakeMonit) StartAndWaitCallCount() int {
	fake.startAndWaitMutex.RLock()
	defer fake.startAndWaitMutex.RUnlock()
	return len(fake.startAndWaitArgsForCall)
}

//StartAndWaitArgsForCall ...
func (fake *FakeMonit) StartAndWaitArgsForCall(i int) string {
	fake.startAndWaitMutex.RLock()
	defer fake.startAndWaitMutex.RUnlock()
	return fake.startAndWaitArgsForCall[i].arg1
}

//StartAndWaitReturns ...
func (fake *FakeMonit) StartAndWaitReturns(result1 error) {
	fake.StartAndWaitStub = nil
	fake.startAndWaitReturns = struct {
		result1 error
	}{result1}
}

//StopAndWait ...
func (fake *FakeMonit) StopAndWait(arg1 string) error {
	fake.stopAndWaitMutex.Lock()
	fake.stopAndWaitArgsForCall = append(fake.stopAndWaitArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("StopAndWait", []interface{}{arg1})
	fake.stopAndWaitMutex.Unlock()
	if fake.StopAndWaitStub != nil {
		return fake.StopAndWaitStub(arg1)
	}
	return fake.stopAndWaitReturns.result1
}

//StopAndWaitCallCount ...
func (fake *FakeMonit) StopAndWaitCallCount() int {
	fake.stopAndWaitMutex.RLock()
	defer fake.stopAndWaitMutex.RUnlock()
	return len(fake.stopAndWaitArgsForCall)
}

//StopAndWaitArgsForCall ...
func (fake *FakeMonit) StopAndWaitArgsForCall(i int) string {
	fake.stopAndWaitMutex.RLock()
	defer fake.stopAndWaitMutex.RUnlock()
	return fake.stopAndWaitArgsForCall[i].arg1
}

//StopAndWaitReturns ...
func (fake *FakeMonit) StopAndWaitReturns(result1 error) {
	fake.StopAndWaitStub = nil
	fake.stopAndWaitReturns = struct {
		result1 error
	}{result1}
}

//SetMonitrcPath ...
func (fake *FakeMonit) SetMonitrcPath(arg1 string) {
	fake.setMonitrcPathMutex.Lock()
	fake.setMonitrcPathArgsForCall = append(fake.setMonitrcPathArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SetMonitrcPath", []interface{}{arg1})
	fake.setMonitrcPathMutex.Unlock()
	if fake.SetMonitrcPathStub != nil {
		fake.SetMonitrcPathStub(arg1)
	}
}

//SetMonitrcPathCallCount ...
func (fake *FakeMonit) SetMonitrcPathCallCount() int {
	fake.setMonitrcPathMutex.RLock()
	defer fake.setMonitrcPathMutex.RUnlock()
	return len(fake.setMonitrcPathArgsForCall)
}

//SetMonitrcPathArgsForCall ...
func (fake *FakeMonit) SetMonitrcPathArgsForCall(i int) string {
	fake.setMonitrcPathMutex.RLock()
	defer fake.setMonitrcPathMutex.RUnlock()
	return fake.setMonitrcPathArgsForCall[i].arg1
}

//SetExecutable ...
func (fake *FakeMonit) SetExecutable(arg1 string) {
	fake.setExecutableMutex.Lock()
	fake.setExecutableArgsForCall = append(fake.setExecutableArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SetExecutable", []interface{}{arg1})
	fake.setExecutableMutex.Unlock()
	if fake.SetExecutableStub != nil {
		fake.SetExecutableStub(arg1)
	}
}

//SetExecutableCallCount ...
func (fake *FakeMonit) SetExecutableCallCount() int {
	fake.setExecutableMutex.RLock()
	defer fake.setExecutableMutex.RUnlock()
	return len(fake.setExecutableArgsForCall)
}

//SetExecutableArgsForCall ...
func (fake *FakeMonit) SetExecutableArgsForCall(i int) string {
	fake.setExecutableMutex.RLock()
	defer fake.setExecutableMutex.RUnlock()
	return fake.setExecutableArgsForCall[i].arg1
}

//Invocations ...
func (fake *FakeMonit) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getSummaryMutex.RLock()
	defer fake.getSummaryMutex.RUnlock()
	fake.getStatusMutex.RLock()
	defer fake.getStatusMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	fake.startAndWaitMutex.RLock()
	defer fake.startAndWaitMutex.RUnlock()
	fake.stopAndWaitMutex.RLock()
	defer fake.stopAndWaitMutex.RUnlock()
	fake.setMonitrcPathMutex.RLock()
	defer fake.setMonitrcPathMutex.RUnlock()
	fake.setExecutableMutex.RLock()
	defer fake.setExecutableMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeMonit) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ monit.Monit = new(FakeMonit)
