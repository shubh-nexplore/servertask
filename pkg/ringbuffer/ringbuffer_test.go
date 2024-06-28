package ringbuffer

import (
	"testing"
	"time"
)

// Test adding timestamps to the ring buffer and counting requests within the last 60 seconds
func TestRingBuffer_AddAndCountRequests(t *testing.T) {
	// Initialize a ring buffer with capacity 5
	rb := NewRingBuffer(5) 

	now := time.Now()

	// Add timestamps to the ring buffer
	rb.Add(now.Add(-70 * time.Second)) // Outside 60-second window
	rb.Add(now.Add(-50 * time.Second)) // Inside 60-second window
	rb.Add(now.Add(-30 * time.Second)) // Inside 60-second window
	rb.Add(now.Add(-10 * time.Second)) // Inside 60-second window
	rb.Add(now)                        // Inside 60-second window

	// Count requests within the last 60 seconds
	count := rb.CountRequestSince(now.Add(-60 * time.Second))
	// Only the last 4 timestamps should be within the window
	expectedCount := 4 

	if count != expectedCount {
		t.Errorf("expected %d, got %d", expectedCount, count)
	}
}

// Test the ring buffer's behavior when it overwrites the oldest timestamps
func TestRingBuffer_OverwriteOldest(t *testing.T) {
	// Initialize a ring buffer with capacity 3
	rb := NewRingBuffer(3) 

	now := time.Now()

	// Add timestamps to the ring buffer
	rb.Add(now.Add(-70 * time.Second)) // This will be overwritten
	rb.Add(now.Add(-50 * time.Second)) // Inside 60-second window
	rb.Add(now.Add(-30 * time.Second)) // Inside 60-second window
	rb.Add(now)                        // This will overwrite the first timestamp

	// Count requests within the last 60 seconds
	count := rb.CountRequestSince(now.Add(-60 * time.Second))
	 // Should count 3 requests, including the one at -50s, -30s, and now
	expectedCount := 3

	if count != expectedCount {
		t.Errorf("expected %d, got %d", expectedCount, count)
	}
}
