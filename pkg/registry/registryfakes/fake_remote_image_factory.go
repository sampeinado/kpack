// Code generated by counterfeiter. DO NOT EDIT.
package registryfakes

import (
	"sync"

	"github.com/pivotal/build-service-system/pkg/registry"
)

type FakeRemoteImageFactory struct {
	NewRemoteStub        func(registry.ImageRef) (registry.RemoteImage, error)
	newRemoteMutex       sync.RWMutex
	newRemoteArgsForCall []struct {
		arg1 registry.ImageRef
	}
	newRemoteReturns struct {
		result1 registry.RemoteImage
		result2 error
	}
	newRemoteReturnsOnCall map[int]struct {
		result1 registry.RemoteImage
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRemoteImageFactory) NewRemote(arg1 registry.ImageRef) (registry.RemoteImage, error) {
	fake.newRemoteMutex.Lock()
	ret, specificReturn := fake.newRemoteReturnsOnCall[len(fake.newRemoteArgsForCall)]
	fake.newRemoteArgsForCall = append(fake.newRemoteArgsForCall, struct {
		arg1 registry.ImageRef
	}{arg1})
	fake.recordInvocation("NewRemote", []interface{}{arg1})
	fake.newRemoteMutex.Unlock()
	if fake.NewRemoteStub != nil {
		return fake.NewRemoteStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.newRemoteReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRemoteImageFactory) NewRemoteCallCount() int {
	fake.newRemoteMutex.RLock()
	defer fake.newRemoteMutex.RUnlock()
	return len(fake.newRemoteArgsForCall)
}

func (fake *FakeRemoteImageFactory) NewRemoteCalls(stub func(registry.ImageRef) (registry.RemoteImage, error)) {
	fake.newRemoteMutex.Lock()
	defer fake.newRemoteMutex.Unlock()
	fake.NewRemoteStub = stub
}

func (fake *FakeRemoteImageFactory) NewRemoteArgsForCall(i int) registry.ImageRef {
	fake.newRemoteMutex.RLock()
	defer fake.newRemoteMutex.RUnlock()
	argsForCall := fake.newRemoteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRemoteImageFactory) NewRemoteReturns(result1 registry.RemoteImage, result2 error) {
	fake.newRemoteMutex.Lock()
	defer fake.newRemoteMutex.Unlock()
	fake.NewRemoteStub = nil
	fake.newRemoteReturns = struct {
		result1 registry.RemoteImage
		result2 error
	}{result1, result2}
}

func (fake *FakeRemoteImageFactory) NewRemoteReturnsOnCall(i int, result1 registry.RemoteImage, result2 error) {
	fake.newRemoteMutex.Lock()
	defer fake.newRemoteMutex.Unlock()
	fake.NewRemoteStub = nil
	if fake.newRemoteReturnsOnCall == nil {
		fake.newRemoteReturnsOnCall = make(map[int]struct {
			result1 registry.RemoteImage
			result2 error
		})
	}
	fake.newRemoteReturnsOnCall[i] = struct {
		result1 registry.RemoteImage
		result2 error
	}{result1, result2}
}

func (fake *FakeRemoteImageFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newRemoteMutex.RLock()
	defer fake.newRemoteMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRemoteImageFactory) recordInvocation(key string, args []interface{}) {
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

var _ registry.RemoteImageFactory = new(FakeRemoteImageFactory)