## Project Structure

- `cmd/` — application entry point  
- `internal/` — core business logic:
  - `logging/` — initialization and configuration of logging
  - `model/` — data structures for use cases, requests, responses, and templates
  - `service/` — request handling, routing, and template rendering
  - `transport/` — HTTP server, handlers, and interfaces
- `templates/` — example templates for mocks
- `testfile.txt` — example test file
- `vendor/` — external dependencies

## Quick Start

1. Install dependencies:
    ```sh
    go mod tidy
    ```

2. Build and run the service:
    ```sh
    go run cmd/gomock.go
    ```

3. Add your templates to the [`templates/`](templates/) folder.

## Usage Example

Once running, the service listens for HTTP requests, matches them to templates, and returns the corresponding responses.

## Testing

To run unit tests, use:
```sh
go test ./...
```
Tests check the core functions of the service and its stability.