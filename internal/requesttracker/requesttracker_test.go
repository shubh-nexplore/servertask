package requesttracker

import (
	"os"
	"testing"
	"time"

	"github.com/insurance/pkg/ringbuffer"
)

const testFilePath = "sample_timestamp_data.json"

// Helper function to set up a RequestTracker for testing
func setup() *RequestTracker {
	_ = os.Remove(testFilePath)
	
	tracker := &RequestTracker{
		requests: ringbuffer.NewRingBuffer(bufferCapacity),
	}
	return tracker
}

// Helper function to clean up after tests
func teardown() {
	_ = os.Remove(testFilePath)
}

// Test recording requests and counting them within the window duration
func TestRequestTracker_RecordAndCountRequests(t *testing.T) {
	tracker := setup()
	// Clean up after test
	defer teardown() 

	// Record requests
	tracker.RecordRequest()
	time.Sleep(1 * time.Second)
	tracker.RecordRequest()
	time.Sleep(1 * time.Second)
	tracker.RecordRequest()

	// Count requests within the last 60 seconds
	count := tracker.CountRequest()
	// Recorded 3 requests
	expectedCount := 3 

	if count != expectedCount {
		t.Errorf("expected %d, got %d", expectedCount, count)
	}
}

// Test persistence by saving and loading request data
func TestRequestTracker_Persistence(t *testing.T) {
	tracker := setup()
	// Clean up after test
	defer teardown() 

	// Record a request and save it
	tracker.RecordRequest()
	tracker.SaveRequestData()

	// Create a new tracker and load the saved data
	newTracker := setup()
	newTracker.loadRequestData()
	count := newTracker.CountRequest()
	// Saved 1 request
	expectedCount := 1

	if count != expectedCount {
		t.Errorf("expected %d, got %d", expectedCount, count)
	}
}
