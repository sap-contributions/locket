// Code generated by counterfeiter. DO NOT EDIT.
package helpersfakes

import (
	"sync"
	"time"

	"github.com/sap-contributions/locket/metrics/helpers"
)

type FakeRequestMetrics struct {
	IncrementRequestsStartedCounterStub        func(requestType string, delta int)
	incrementRequestsStartedCounterMutex       sync.RWMutex
	incrementRequestsStartedCounterArgsForCall []struct {
		requestType string
		delta       int
	}
	IncrementRequestsSucceededCounterStub        func(requestType string, delta int)
	incrementRequestsSucceededCounterMutex       sync.RWMutex
	incrementRequestsSucceededCounterArgsForCall []struct {
		requestType string
		delta       int
	}
	IncrementRequestsFailedCounterStub        func(requestType string, delta int)
	incrementRequestsFailedCounterMutex       sync.RWMutex
	incrementRequestsFailedCounterArgsForCall []struct {
		requestType string
		delta       int
	}
	IncrementRequestsInFlightCounterStub        func(requestType string, delta int)
	incrementRequestsInFlightCounterMutex       sync.RWMutex
	incrementRequestsInFlightCounterArgsForCall []struct {
		requestType string
		delta       int
	}
	DecrementRequestsInFlightCounterStub        func(requestType string, delta int)
	decrementRequestsInFlightCounterMutex       sync.RWMutex
	decrementRequestsInFlightCounterArgsForCall []struct {
		requestType string
		delta       int
	}
	IncrementRequestsCancelledCounterStub        func(requestType string, delta int)
	incrementRequestsCancelledCounterMutex       sync.RWMutex
	incrementRequestsCancelledCounterArgsForCall []struct {
		requestType string
		delta       int
	}
	UpdateLatencyStub        func(requestType string, dur time.Duration)
	updateLatencyMutex       sync.RWMutex
	updateLatencyArgsForCall []struct {
		requestType string
		dur         time.Duration
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRequestMetrics) IncrementRequestsStartedCounter(requestType string, delta int) {
	fake.incrementRequestsStartedCounterMutex.Lock()
	fake.incrementRequestsStartedCounterArgsForCall = append(fake.incrementRequestsStartedCounterArgsForCall, struct {
		requestType string
		delta       int
	}{requestType, delta})
	fake.recordInvocation("IncrementRequestsStartedCounter", []interface{}{requestType, delta})
	fake.incrementRequestsStartedCounterMutex.Unlock()
	if fake.IncrementRequestsStartedCounterStub != nil {
		fake.IncrementRequestsStartedCounterStub(requestType, delta)
	}
}

func (fake *FakeRequestMetrics) IncrementRequestsStartedCounterCallCount() int {
	fake.incrementRequestsStartedCounterMutex.RLock()
	defer fake.incrementRequestsStartedCounterMutex.RUnlock()
	return len(fake.incrementRequestsStartedCounterArgsForCall)
}

func (fake *FakeRequestMetrics) IncrementRequestsStartedCounterArgsForCall(i int) (string, int) {
	fake.incrementRequestsStartedCounterMutex.RLock()
	defer fake.incrementRequestsStartedCounterMutex.RUnlock()
	return fake.incrementRequestsStartedCounterArgsForCall[i].requestType, fake.incrementRequestsStartedCounterArgsForCall[i].delta
}

func (fake *FakeRequestMetrics) IncrementRequestsSucceededCounter(requestType string, delta int) {
	fake.incrementRequestsSucceededCounterMutex.Lock()
	fake.incrementRequestsSucceededCounterArgsForCall = append(fake.incrementRequestsSucceededCounterArgsForCall, struct {
		requestType string
		delta       int
	}{requestType, delta})
	fake.recordInvocation("IncrementRequestsSucceededCounter", []interface{}{requestType, delta})
	fake.incrementRequestsSucceededCounterMutex.Unlock()
	if fake.IncrementRequestsSucceededCounterStub != nil {
		fake.IncrementRequestsSucceededCounterStub(requestType, delta)
	}
}

func (fake *FakeRequestMetrics) IncrementRequestsSucceededCounterCallCount() int {
	fake.incrementRequestsSucceededCounterMutex.RLock()
	defer fake.incrementRequestsSucceededCounterMutex.RUnlock()
	return len(fake.incrementRequestsSucceededCounterArgsForCall)
}

func (fake *FakeRequestMetrics) IncrementRequestsSucceededCounterArgsForCall(i int) (string, int) {
	fake.incrementRequestsSucceededCounterMutex.RLock()
	defer fake.incrementRequestsSucceededCounterMutex.RUnlock()
	return fake.incrementRequestsSucceededCounterArgsForCall[i].requestType, fake.incrementRequestsSucceededCounterArgsForCall[i].delta
}

func (fake *FakeRequestMetrics) IncrementRequestsFailedCounter(requestType string, delta int) {
	fake.incrementRequestsFailedCounterMutex.Lock()
	fake.incrementRequestsFailedCounterArgsForCall = append(fake.incrementRequestsFailedCounterArgsForCall, struct {
		requestType string
		delta       int
	}{requestType, delta})
	fake.recordInvocation("IncrementRequestsFailedCounter", []interface{}{requestType, delta})
	fake.incrementRequestsFailedCounterMutex.Unlock()
	if fake.IncrementRequestsFailedCounterStub != nil {
		fake.IncrementRequestsFailedCounterStub(requestType, delta)
	}
}

func (fake *FakeRequestMetrics) IncrementRequestsFailedCounterCallCount() int {
	fake.incrementRequestsFailedCounterMutex.RLock()
	defer fake.incrementRequestsFailedCounterMutex.RUnlock()
	return len(fake.incrementRequestsFailedCounterArgsForCall)
}

func (fake *FakeRequestMetrics) IncrementRequestsFailedCounterArgsForCall(i int) (string, int) {
	fake.incrementRequestsFailedCounterMutex.RLock()
	defer fake.incrementRequestsFailedCounterMutex.RUnlock()
	return fake.incrementRequestsFailedCounterArgsForCall[i].requestType, fake.incrementRequestsFailedCounterArgsForCall[i].delta
}

func (fake *FakeRequestMetrics) IncrementRequestsInFlightCounter(requestType string, delta int) {
	fake.incrementRequestsInFlightCounterMutex.Lock()
	fake.incrementRequestsInFlightCounterArgsForCall = append(fake.incrementRequestsInFlightCounterArgsForCall, struct {
		requestType string
		delta       int
	}{requestType, delta})
	fake.recordInvocation("IncrementRequestsInFlightCounter", []interface{}{requestType, delta})
	fake.incrementRequestsInFlightCounterMutex.Unlock()
	if fake.IncrementRequestsInFlightCounterStub != nil {
		fake.IncrementRequestsInFlightCounterStub(requestType, delta)
	}
}

func (fake *FakeRequestMetrics) IncrementRequestsInFlightCounterCallCount() int {
	fake.incrementRequestsInFlightCounterMutex.RLock()
	defer fake.incrementRequestsInFlightCounterMutex.RUnlock()
	return len(fake.incrementRequestsInFlightCounterArgsForCall)
}

func (fake *FakeRequestMetrics) IncrementRequestsInFlightCounterArgsForCall(i int) (string, int) {
	fake.incrementRequestsInFlightCounterMutex.RLock()
	defer fake.incrementRequestsInFlightCounterMutex.RUnlock()
	return fake.incrementRequestsInFlightCounterArgsForCall[i].requestType, fake.incrementRequestsInFlightCounterArgsForCall[i].delta
}

func (fake *FakeRequestMetrics) DecrementRequestsInFlightCounter(requestType string, delta int) {
	fake.decrementRequestsInFlightCounterMutex.Lock()
	fake.decrementRequestsInFlightCounterArgsForCall = append(fake.decrementRequestsInFlightCounterArgsForCall, struct {
		requestType string
		delta       int
	}{requestType, delta})
	fake.recordInvocation("DecrementRequestsInFlightCounter", []interface{}{requestType, delta})
	fake.decrementRequestsInFlightCounterMutex.Unlock()
	if fake.DecrementRequestsInFlightCounterStub != nil {
		fake.DecrementRequestsInFlightCounterStub(requestType, delta)
	}
}

func (fake *FakeRequestMetrics) DecrementRequestsInFlightCounterCallCount() int {
	fake.decrementRequestsInFlightCounterMutex.RLock()
	defer fake.decrementRequestsInFlightCounterMutex.RUnlock()
	return len(fake.decrementRequestsInFlightCounterArgsForCall)
}

func (fake *FakeRequestMetrics) DecrementRequestsInFlightCounterArgsForCall(i int) (string, int) {
	fake.decrementRequestsInFlightCounterMutex.RLock()
	defer fake.decrementRequestsInFlightCounterMutex.RUnlock()
	return fake.decrementRequestsInFlightCounterArgsForCall[i].requestType, fake.decrementRequestsInFlightCounterArgsForCall[i].delta
}

func (fake *FakeRequestMetrics) IncrementRequestsCancelledCounter(requestType string, delta int) {
	fake.incrementRequestsCancelledCounterMutex.Lock()
	fake.incrementRequestsCancelledCounterArgsForCall = append(fake.incrementRequestsCancelledCounterArgsForCall, struct {
		requestType string
		delta       int
	}{requestType, delta})
	fake.recordInvocation("IncrementRequestsCancelledCounter", []interface{}{requestType, delta})
	fake.incrementRequestsCancelledCounterMutex.Unlock()
	if fake.IncrementRequestsCancelledCounterStub != nil {
		fake.IncrementRequestsCancelledCounterStub(requestType, delta)
	}
}

func (fake *FakeRequestMetrics) IncrementRequestsCancelledCounterCallCount() int {
	fake.incrementRequestsCancelledCounterMutex.RLock()
	defer fake.incrementRequestsCancelledCounterMutex.RUnlock()
	return len(fake.incrementRequestsCancelledCounterArgsForCall)
}

func (fake *FakeRequestMetrics) IncrementRequestsCancelledCounterArgsForCall(i int) (string, int) {
	fake.incrementRequestsCancelledCounterMutex.RLock()
	defer fake.incrementRequestsCancelledCounterMutex.RUnlock()
	return fake.incrementRequestsCancelledCounterArgsForCall[i].requestType, fake.incrementRequestsCancelledCounterArgsForCall[i].delta
}

func (fake *FakeRequestMetrics) UpdateLatency(requestType string, dur time.Duration) {
	fake.updateLatencyMutex.Lock()
	fake.updateLatencyArgsForCall = append(fake.updateLatencyArgsForCall, struct {
		requestType string
		dur         time.Duration
	}{requestType, dur})
	fake.recordInvocation("UpdateLatency", []interface{}{requestType, dur})
	fake.updateLatencyMutex.Unlock()
	if fake.UpdateLatencyStub != nil {
		fake.UpdateLatencyStub(requestType, dur)
	}
}

func (fake *FakeRequestMetrics) UpdateLatencyCallCount() int {
	fake.updateLatencyMutex.RLock()
	defer fake.updateLatencyMutex.RUnlock()
	return len(fake.updateLatencyArgsForCall)
}

func (fake *FakeRequestMetrics) UpdateLatencyArgsForCall(i int) (string, time.Duration) {
	fake.updateLatencyMutex.RLock()
	defer fake.updateLatencyMutex.RUnlock()
	return fake.updateLatencyArgsForCall[i].requestType, fake.updateLatencyArgsForCall[i].dur
}

func (fake *FakeRequestMetrics) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.incrementRequestsStartedCounterMutex.RLock()
	defer fake.incrementRequestsStartedCounterMutex.RUnlock()
	fake.incrementRequestsSucceededCounterMutex.RLock()
	defer fake.incrementRequestsSucceededCounterMutex.RUnlock()
	fake.incrementRequestsFailedCounterMutex.RLock()
	defer fake.incrementRequestsFailedCounterMutex.RUnlock()
	fake.incrementRequestsInFlightCounterMutex.RLock()
	defer fake.incrementRequestsInFlightCounterMutex.RUnlock()
	fake.decrementRequestsInFlightCounterMutex.RLock()
	defer fake.decrementRequestsInFlightCounterMutex.RUnlock()
	fake.incrementRequestsCancelledCounterMutex.RLock()
	defer fake.incrementRequestsCancelledCounterMutex.RUnlock()
	fake.updateLatencyMutex.RLock()
	defer fake.updateLatencyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRequestMetrics) recordInvocation(key string, args []interface{}) {
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

var _ helpers.RequestMetrics = new(FakeRequestMetrics)
