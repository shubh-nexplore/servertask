# servertask
Using only the standard library, create a Go HTTP server that on each request responds with a counter of the total number of requests that it has received during the previous 60 seconds (moving window). The server should continue to return the correct numbers after restarting it, by persisting data to a file.


## Setup

1. **Clone the repository**:

    ```sh
    git clone https://github.com/shubh-nexplore/servertask.git
    cd myapp
    ```

2. **Install dependencies**:

    Ensure you have [Go](https://golang.org/dl/) installed.
   
## Running the Application

To run the application:

```sh
go run cmd/app/main.go
```

The application will start and listen on port 8080. We can access it at `http://localhost:8080`.

Testing
To run the tests:

```
go test ./...
```

## Project Details

1. **HTTP Server**:
    - Utilized the `net/http` package to create a basic HTTP server.
    - Defined a handler that processes incoming requests.

2. **Tracking Requests**:
    - To efficiently manage and store timestamps of incoming requests, I implemented a ring buffer (or circular buffer). A ring buffer is a fixed-size data structure that uses a single, continuous buffer and wraps around when it reaches the end. This ensures efficient use of memory and constant time complexity for adding and removing elements.The ring buffer uses a fixed amount of memory, avoiding the overhead and complexity of dynamic resizing. Both insertion and removal operations are O(1), ensuring predictable and constant performance even under moderate to high loads.
    - The ring buffer allows us to maintain a sliding window of timestamps, which is essential for counting requests within the last 60 seconds.

3. **Moving Window**:
    - Each request is recorded with its timestamp.
    - For each request, the handler calculates the number of requests in the last 60 seconds by removing outdated timestamps from the buffer.

4. **Persistence**:
    - Implemented methods to load and save request data from/to a JSON file.
    - Ensured the server loads the data on startup to maintain continuity across restarts. This allows the application to continue returning accurate request counts even after a restart.

5. **Graceful Shutdown**:
    - Added signal handling to save the current state before shutting down. This ensures no data loss occurs when the server is stopped.
    - Periodically saves the state to minimize data loss in case of unexpected shutdowns.


## Project Structure

`cmd/myapp/main.go`

The entry point of the application. It initializes the request tracker, sets up signal handling for graceful shutdown, and starts the HTTP server.

`internal/handler/handler.go`

Contains the HTTP handler function that records incoming requests and responds with the number of requests received in the last 60 seconds.

`internal/requesttracker/requesttracker.go`

Manages the ring buffer and handles concurrency for tracking requests. It also loads and saves request data to a file.

`pkg/ringbuffer/ringbuffer.go`

Implements the ring buffer that efficiently manages request timestamps.
