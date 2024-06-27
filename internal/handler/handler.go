package handler

import (
	"fmt"
	"net/http"

	"github.com/insurance/internal/requesttracker"
)

// NewHandler returns an HTTP handler function that records incoming requests
// and responds with the number of requests received in the last 60 seconds.
func NewHandler(tracker *requesttracker.RequestTracker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Record the incoming request
		tracker.RecordRequest()

		// Count the number of requests in the last 60 seconds
		requestCount := tracker.CountRequest()

		// Respond with the request count
		fmt.Fprintf(w, "Requests in the last 60 seconds: %d", requestCount)
	}
}

