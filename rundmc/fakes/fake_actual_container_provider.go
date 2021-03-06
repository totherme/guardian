// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/guardian/rundmc"
)

type FakeActualContainerProvider struct {
	ProvideStub        func(directory string) (rundmc.ActualContainer, error)
	provideMutex       sync.RWMutex
	provideArgsForCall []struct {
		directory string
	}
	provideReturns struct {
		result1 rundmc.ActualContainer
		result2 error
	}
}

func (fake *FakeActualContainerProvider) Provide(directory string) (rundmc.ActualContainer, error) {
	fake.provideMutex.Lock()
	fake.provideArgsForCall = append(fake.provideArgsForCall, struct {
		directory string
	}{directory})
	fake.provideMutex.Unlock()
	if fake.ProvideStub != nil {
		return fake.ProvideStub(directory)
	} else {
		return fake.provideReturns.result1, fake.provideReturns.result2
	}
}

func (fake *FakeActualContainerProvider) ProvideCallCount() int {
	fake.provideMutex.RLock()
	defer fake.provideMutex.RUnlock()
	return len(fake.provideArgsForCall)
}

func (fake *FakeActualContainerProvider) ProvideArgsForCall(i int) string {
	fake.provideMutex.RLock()
	defer fake.provideMutex.RUnlock()
	return fake.provideArgsForCall[i].directory
}

func (fake *FakeActualContainerProvider) ProvideReturns(result1 rundmc.ActualContainer, result2 error) {
	fake.ProvideStub = nil
	fake.provideReturns = struct {
		result1 rundmc.ActualContainer
		result2 error
	}{result1, result2}
}

var _ rundmc.ActualContainerProvider = new(FakeActualContainerProvider)
