# README.md

# SSE Scenario Server

This project is a Go application that serves different scenarios to clients through Server-Sent Events (SSE). Scenarios are accessible via the endpoint `/tests/scenarios/:scenario`.

## Project Structure

```
sse-scenario-server
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── handlers
│   │   ├── scenario.go  # Handler for retrieving scenarios
│   │   └── sse.go       # Handler for streaming SSE events
│   ├── models
│   │   └── scenario.go   # Definition of the Scenario struct
│   └── server
│       └── server.go    # Server setup and route registration
├── tests
│   └── scenarios
│       └── sample.json   # Sample scenario data for testing
├── go.mod                # Module definition and dependencies
├── go.sum                # Checksums for module dependencies
└── README.md             # Project documentation
```

## Setup Instructions

1. Clone the repository:
   ```
   git clone <repository-url>
   cd sse-scenario-server
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Run the application:
   ```
   go run cmd/main.go
   ```

## Building and Running

To build the application, run the following command:
```
# Build the image
docker build -t hockeykit-tester .

# Run with docker-compose
docker-compose up
```

## Usage

Once the server is running, you can access scenarios through the following endpoint:
```
GET /tests/scenarios/:scenario
```

Replace `:scenario` with the ID of the scenario you wish to retrieve. The server will stream the scenario data to the client using SSE.