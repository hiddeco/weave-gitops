// Code generated by counterfeiter. DO NOT EDIT.
package cachefakes

import (
	"context"
	"sync"

	"github.com/weaveworks/weave-gitops/core/cache"
)

type FakeContainer struct {
	ForceRefreshStub        func(cache.StorageType) error
	forceRefreshMutex       sync.RWMutex
	forceRefreshArgsForCall []struct {
		arg1 cache.StorageType
	}
	forceRefreshReturns struct {
		result1 error
	}
	forceRefreshReturnsOnCall map[int]struct {
		result1 error
	}
	ListStub        func(cache.StorageType) (interface{}, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 cache.StorageType
	}
	listReturns struct {
		result1 interface{}
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 interface{}
		result2 error
	}
	StartStub        func(context.Context)
	startMutex       sync.RWMutex
	startArgsForCall []struct {
		arg1 context.Context
	}
	StopStub        func()
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeContainer) ForceRefresh(arg1 cache.StorageType) error {
	fake.forceRefreshMutex.Lock()
	ret, specificReturn := fake.forceRefreshReturnsOnCall[len(fake.forceRefreshArgsForCall)]
	fake.forceRefreshArgsForCall = append(fake.forceRefreshArgsForCall, struct {
		arg1 cache.StorageType
	}{arg1})
	stub := fake.ForceRefreshStub
	fakeReturns := fake.forceRefreshReturns
	fake.recordInvocation("ForceRefresh", []interface{}{arg1})
	fake.forceRefreshMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeContainer) ForceRefreshCallCount() int {
	fake.forceRefreshMutex.RLock()
	defer fake.forceRefreshMutex.RUnlock()
	return len(fake.forceRefreshArgsForCall)
}

func (fake *FakeContainer) ForceRefreshCalls(stub func(cache.StorageType) error) {
	fake.forceRefreshMutex.Lock()
	defer fake.forceRefreshMutex.Unlock()
	fake.ForceRefreshStub = stub
}

func (fake *FakeContainer) ForceRefreshArgsForCall(i int) cache.StorageType {
	fake.forceRefreshMutex.RLock()
	defer fake.forceRefreshMutex.RUnlock()
	argsForCall := fake.forceRefreshArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeContainer) ForceRefreshReturns(result1 error) {
	fake.forceRefreshMutex.Lock()
	defer fake.forceRefreshMutex.Unlock()
	fake.ForceRefreshStub = nil
	fake.forceRefreshReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) ForceRefreshReturnsOnCall(i int, result1 error) {
	fake.forceRefreshMutex.Lock()
	defer fake.forceRefreshMutex.Unlock()
	fake.ForceRefreshStub = nil
	if fake.forceRefreshReturnsOnCall == nil {
		fake.forceRefreshReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.forceRefreshReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) List(arg1 cache.StorageType) (interface{}, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 cache.StorageType
	}{arg1})
	stub := fake.ListStub
	fakeReturns := fake.listReturns
	fake.recordInvocation("List", []interface{}{arg1})
	fake.listMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeContainer) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeContainer) ListCalls(stub func(cache.StorageType) (interface{}, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeContainer) ListArgsForCall(i int) cache.StorageType {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeContainer) ListReturns(result1 interface{}, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) ListReturnsOnCall(i int, result1 interface{}, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) Start(arg1 context.Context) {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.StartStub
	fake.recordInvocation("Start", []interface{}{arg1})
	fake.startMutex.Unlock()
	if stub != nil {
		fake.StartStub(arg1)
	}
}

func (fake *FakeContainer) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeContainer) StartCalls(stub func(context.Context)) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = stub
}

func (fake *FakeContainer) StartArgsForCall(i int) context.Context {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	argsForCall := fake.startArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeContainer) Stop() {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
	}{})
	stub := fake.StopStub
	fake.recordInvocation("Stop", []interface{}{})
	fake.stopMutex.Unlock()
	if stub != nil {
		fake.StopStub()
	}
}

func (fake *FakeContainer) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeContainer) StopCalls(stub func()) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = stub
}

func (fake *FakeContainer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.forceRefreshMutex.RLock()
	defer fake.forceRefreshMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeContainer) recordInvocation(key string, args []interface{}) {
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

var _ cache.Container = new(FakeContainer)
