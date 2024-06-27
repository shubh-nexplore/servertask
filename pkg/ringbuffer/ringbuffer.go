package ringbuffer

import (
	"time"
)

// RingBuffer efficiently manages request timestamps
type RingBuffer struct {
	Timestamps []time.Time // Slice to store timestamps
	Capacity   int         // Maximum size of the ring buffer
	StartIndex int         // Index of the oldest element
	Count      int         // Current number of elements
}

// NewRingBuffer initializes a new RingBuffer with the provided capacity
func NewRingBuffer(capacity int) *RingBuffer {
	return &RingBuffer{
		Timestamps: make([]time.Time, capacity),
		Capacity:   capacity,
	}
}

// Add inserts a new timestamp into the ring buffer
func (rb *RingBuffer) Add(timestamp time.Time) {
	if rb.Count < rb.Capacity {
		rb.Timestamps[(rb.StartIndex+rb.Count)%rb.Capacity] = timestamp
		rb.Count++
	} else {
		rb.Timestamps[rb.StartIndex] = timestamp
		rb.StartIndex = (rb.StartIndex + 1) % rb.Capacity
	}
}

// CountRequestSince returns the number of timestamps in the ring buffer that are after the given cutoff time
func (rb *RingBuffer) CountRequestSince(cutoffTimestamp time.Time) int {
	count := 0
	for i := 0; i < rb.Count; i++ {
		idx := (rb.StartIndex + i) % rb.Capacity
		if rb.Timestamps[idx].After(cutoffTimestamp) {
			count++
		}
	}
	return count
}
