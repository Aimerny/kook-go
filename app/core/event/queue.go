package event

import (
	"math"
	"sync"

	"github.com/aimerny/kook-go/app/core/model"
)

type EventQueue struct {
	events *sync.Map
	MinSn  int
	mu     *sync.RWMutex
}

func NewEventQueue() *EventQueue {
	return &EventQueue{
		events: &sync.Map{},
		MinSn:  math.MaxInt,
		mu:     &sync.RWMutex{},
	}
}

// Push returns min sn
func (eq *EventQueue) Push(signal *model.Signal) int {
	eq.mu.Lock()
	defer eq.mu.Unlock()
	eq.events.Store(signal.SerialNumber, signal)
	if signal.SerialNumber < eq.MinSn {
		eq.MinSn = signal.SerialNumber
	}
	return eq.MinSn
}

func (eq *EventQueue) Pop() *model.Signal {
	eq.mu.Lock()
	defer eq.mu.Unlock()
	if signal, ok := eq.events.Load(eq.MinSn); ok {
		if signal, ok := signal.(*model.Signal); ok {
			eq.events.Delete(eq.MinSn)
			eq.resetMinKey()
			return signal
		}
	}
	return nil
}

func (eq *EventQueue) IsEmpty() bool {
	empty := true
	eq.events.Range(func(key, value interface{}) bool {
		empty = false
		return false
	})
	return empty
}

func (eq *EventQueue) Clear() {
	eq.mu.Lock()
	defer eq.mu.Unlock()
	eq.events.Range(func(k, v interface{}) bool {
		eq.events.Delete(k)
		return true
	})
	eq.MinSn = math.MaxInt
}

func (eq *EventQueue) resetMinKey() {

	eq.MinSn = math.MaxInt
	eq.events.Range(func(k, v interface{}) bool {
		if k, ok := k.(int); ok {
			if k < eq.MinSn {
				eq.MinSn = k
			}
		}
		return true
	})

}
