package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/insurance/internal/handler"
	"github.com/insurance/internal/requesttracker"
)

func main() {
	// Initialize the request tracker
	tracker := requesttracker.NewRequestTracker()

	// Create a channel to listen for OS signals for graceful shutdown
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	// Goroutine to handle graceful shutdown and save state before exit
	go func() {
		<-signalChannel
		fmt.Println("Exiting ...")
		tracker.SaveRequestData()
		os.Exit(0)
	}()

	// Set up the HTTP handler
	http.HandleFunc("/", handler.NewHandler(tracker))

	// Goroutine to periodically save the tracker state
	go func() {
		for {
			time.Sleep(time.Minute)
			tracker.SaveRequestData()
		}
	}()

	// Start the HTTP server
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

