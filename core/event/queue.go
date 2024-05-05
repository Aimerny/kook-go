package event

import (
	"github.com/aimerny/kook-sdk/core/model"
	"math"
)

type EventQueue struct {
	events map[int]*model.Signal
	MinSn  int
}

func NewEventQueue() *EventQueue {
	return &EventQueue{
		events: make(map[int]*model.Signal),
		MinSn:  math.MaxInt,
	}
}

// Push returns min sn
func (eq *EventQueue) Push(signal *model.Signal) int {
	eq.events[signal.SerialNumber] = signal
	if signal.SerialNumber < eq.MinSn {
		eq.MinSn = signal.SerialNumber
	}
	return eq.MinSn
}

func (eq *EventQueue) Pop() *model.Signal {
	if signal, ok := eq.events[eq.MinSn]; ok {
		eq.events[eq.MinSn] = nil
		delete(eq.events, eq.MinSn)
		eq.resetMinKey()
		return signal
	}
	return nil
}

func (eq *EventQueue) IsEmpty() bool {
	return len(eq.events) == 0
}

func (eq *EventQueue) Clear() {
	eq.events = make(map[int]*model.Signal)
	eq.MinSn = math.MaxInt
}

func (eq *EventQueue) resetMinKey() {
	eq.MinSn = math.MaxInt
	for k := range eq.events {
		if k < eq.MinSn {
			eq.MinSn = k
		}
	}
}
