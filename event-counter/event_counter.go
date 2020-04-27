package eventcounter

import (
	"fmt"
	"sync"
	"time"
)

// EventCounter is a basic struct to keep track of events.
type EventCounter struct {
	Mutex      sync.RWMutex
	Name       string
	Value      uint64
	LastUpdate time.Time
}

// Inc add counter value by 1
func (counter *EventCounter) Inc() {
	counter.Mutex.RLock()
	fmt.Println(counter.Value)
	counter.Value++
	counter.LastUpdate = time.Now()
	counter.Mutex.RUnlock()
	return
}

// Dec add counter value by 1
func (counter *EventCounter) Dec() {
	counter.Mutex.RLock()
	fmt.Println(counter.Value)
	counter.Value--
	counter.LastUpdate = time.Now()
	counter.Mutex.RUnlock()
	return
}

// IncBy add counter value by 1
func (counter *EventCounter) IncBy(num uint64) {
	counter.Mutex.RLock()
	fmt.Println(counter.Value)
	counter.Value = counter.Value + num
	counter.LastUpdate = time.Now()
	counter.Mutex.RUnlock()
	return
}

// DecBy add counter value by 1
func (counter *EventCounter) DecBy(num uint64) {
	counter.Mutex.RLock()
	fmt.Println(counter.Value)
	counter.Value = counter.Value - num
	counter.LastUpdate = time.Now()
	counter.Mutex.RUnlock()
	return
}

// Reset add counter value by 1
func (counter *EventCounter) Reset() {
	counter.Mutex.RLock()
	fmt.Println(counter.Value)
	counter.Value = 0
	counter.LastUpdate = time.Now()
	counter.Mutex.RUnlock()
	return
}
