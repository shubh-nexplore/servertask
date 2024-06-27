package main

import (
	"fmt"
	"net/http"

)

func main() {



	// Start the HTTP server
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

