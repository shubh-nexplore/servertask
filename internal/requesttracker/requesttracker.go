package requesttracker

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/insurance/pkg/ringbuffer"
)

const (
	dataFilePath    = "timestamp_data.json"
	windowDuration  = 60 * time.Second
	bufferCapacity  = 1000
)

// RequestTracker manages the ring buffer and handles concurrency for tracking requests
type RequestTracker struct {
	mu       sync.RWMutex
	requests *ringbuffer.RingBuffer
}

// NewRequestTracker initializes a new RequestTracker and loads the saved request data
func NewRequestTracker() *RequestTracker {
	tracker := &RequestTracker{
		requests: ringbuffer.NewRingBuffer(bufferCapacity),
	}
	tracker.loadRequestData()
	return tracker
}

// loadRequestData reads the saved request timestamps from a file and loads them into the ring buffer
func (tracker *RequestTracker) loadRequestData() {
	file, err := os.Open(dataFilePath)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println("Error opening data file:", err)
		}
		return
	}
	defer file.Close()

	var savedTimestamps []time.Time
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&savedTimestamps)
	if err != nil {
		fmt.Println("Error decoding JSON data:", err)
		return
	}

	tracker.mu.Lock()
	defer tracker.mu.Unlock()
	for _, timestamp := range savedTimestamps {
		tracker.requests.Add(timestamp)
	}
}

// SaveRequestData writes the current request timestamps to a file
func (tracker *RequestTracker) SaveRequestData() {
	tracker.mu.RLock()
	defer tracker.mu.RUnlock()

	file, err := os.Create(dataFilePath)
	if err != nil {
		fmt.Println("Error creating data file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tracker.requests.Timestamps[:tracker.requests.Count])
	if err != nil {
		fmt.Println("Error encoding JSON data:", err)
	}
}

// RecordRequest records a new request timestamp
func (tracker *RequestTracker) RecordRequest() {
	tracker.mu.Lock()
	defer tracker.mu.Unlock()
	tracker.requests.Add(time.Now())
}

// CountRequest returns the number of requests received in the last window duration (60 seconds)
func (tracker *RequestTracker) CountRequest() int {
	tracker.mu.RLock()
	defer tracker.mu.RUnlock()
	cutoffTimestamp := time.Now().Add(-windowDuration)
	return tracker.requests.CountRequestSince(cutoffTimestamp)
}
