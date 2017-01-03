// This file was generated by counterfeiter
// counterfeiter -o iredis/client_fake.go --fake-name ClientFake iredis/client.go Client

package iredis

import "sync"

//ClientFake ...
type ClientFake struct {
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	PingStub        func() StatusCmd
	pingMutex       sync.RWMutex
	pingArgsForCall []struct{}
	pingReturns     struct {
		result1 StatusCmd
	}
	BgRewriteAOFStub        func() StatusCmd
	bgRewriteAOFMutex       sync.RWMutex
	bgRewriteAOFArgsForCall []struct{}
	bgRewriteAOFReturns     struct {
		result1 StatusCmd
	}
	InfoStub        func(...string) StringCmd
	infoMutex       sync.RWMutex
	infoArgsForCall []struct {
		arg1 []string
	}
	infoReturns struct {
		result1 StringCmd
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

//Close ...
func (fake *ClientFake) Close() error {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	return fake.closeReturns.result1
}

//CloseCallCount ...
func (fake *ClientFake) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

//CloseReturns ...
func (fake *ClientFake) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

//Ping ...
func (fake *ClientFake) Ping() StatusCmd {
	fake.pingMutex.Lock()
	fake.pingArgsForCall = append(fake.pingArgsForCall, struct{}{})
	fake.recordInvocation("Ping", []interface{}{})
	fake.pingMutex.Unlock()
	if fake.PingStub != nil {
		return fake.PingStub()
	}
	return fake.pingReturns.result1
}

//PingCallCount ...
func (fake *ClientFake) PingCallCount() int {
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	return len(fake.pingArgsForCall)
}

//PingReturns ...
func (fake *ClientFake) PingReturns(result1 StatusCmd) {
	fake.PingStub = nil
	fake.pingReturns = struct {
		result1 StatusCmd
	}{result1}
}

//BgRewriteAOF ...
func (fake *ClientFake) BgRewriteAOF() StatusCmd {
	fake.bgRewriteAOFMutex.Lock()
	fake.bgRewriteAOFArgsForCall = append(fake.bgRewriteAOFArgsForCall, struct{}{})
	fake.recordInvocation("BgRewriteAOF", []interface{}{})
	fake.bgRewriteAOFMutex.Unlock()
	if fake.BgRewriteAOFStub != nil {
		return fake.BgRewriteAOFStub()
	}
	return fake.bgRewriteAOFReturns.result1
}

//BgRewriteAOFCallCount ...
func (fake *ClientFake) BgRewriteAOFCallCount() int {
	fake.bgRewriteAOFMutex.RLock()
	defer fake.bgRewriteAOFMutex.RUnlock()
	return len(fake.bgRewriteAOFArgsForCall)
}

//BgRewriteAOFReturns ...
func (fake *ClientFake) BgRewriteAOFReturns(result1 StatusCmd) {
	fake.BgRewriteAOFStub = nil
	fake.bgRewriteAOFReturns = struct {
		result1 StatusCmd
	}{result1}
}

//Info ...
func (fake *ClientFake) Info(arg1 ...string) StringCmd {
	fake.infoMutex.Lock()
	fake.infoArgsForCall = append(fake.infoArgsForCall, struct {
		arg1 []string
	}{arg1})
	fake.recordInvocation("Info", []interface{}{arg1})
	fake.infoMutex.Unlock()
	if fake.InfoStub != nil {
		return fake.InfoStub(arg1...)
	}
	return fake.infoReturns.result1
}

//InfoCallCount ...
func (fake *ClientFake) InfoCallCount() int {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return len(fake.infoArgsForCall)
}

//InfoArgsForCall ...
func (fake *ClientFake) InfoArgsForCall(i int) []string {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return fake.infoArgsForCall[i].arg1
}

//InfoReturns ...
func (fake *ClientFake) InfoReturns(result1 StringCmd) {
	fake.InfoStub = nil
	fake.infoReturns = struct {
		result1 StringCmd
	}{result1}
}

//Invocations ...
func (fake *ClientFake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	fake.bgRewriteAOFMutex.RLock()
	defer fake.bgRewriteAOFMutex.RUnlock()
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return fake.invocations
}

func (fake *ClientFake) recordInvocation(key string, args []interface{}) {
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

var _ Client = new(ClientFake)
