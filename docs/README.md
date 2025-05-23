## Project Structure

- `cmd/` — application entry point  
- `internal/` — core business logic:
  - `logging/` — initialization and configuration of logging
  - `model/` — data structures for use cases, requests, responses, and templates
  - `service/` — request handling, routing, and template rendering
    - `serivice/builder` - route builder, template builder
    - `service/matcher` - request matcher
    - `service/preparer` - response preparer
  - `transport/` — HTTP server, handlers, and interfaces
- `vendor/` — external dependencies

## Quick Start

1. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Add folder with templates.

3. Build and run the service:
    ```sh
    go run cmd/gomock.go  --dir /path/to/templates/folder --address :8080
    ```


## Usage Example

Once running, the service listens for HTTP requests, matches them to templates, and returns the corresponding responses.

## Testing

To run unit tests, use:
```sh
go test ./...
```
Tests check the core functions of the service and its stability.

## Syntax 
### Path Parameters
- `:param_name` - path parameter that can be matched with any value
- `{id:[a-zA-Z0-9-]+}` - path parameter that can be matched with a regular expression

### Allowed values
- `GET` - HTTP method
- `POST` - HTTP method
- `PUT` - HTTP method
- `DELETE` - HTTP method
- `PATCH` - HTTP method
- `OPTIONS` - HTTP method
- `HEAD` - HTTP method
- `TRACE` - HTTP method
- `CONNECT` - HTTP method

### Placeholder Syntax
- `${...}` - any value
- `${regexp:...}` - value that matches the regular expression
- `${req.queryParams:...}` - value from query parameters
- `${req.pathParams:...}` - value from path parameters
- `${req.formParams:...}` - value from form parameters
- `${req.headers:...}` - value from headers
- `${req.body:...}` - value from body

### Requst Matching
- `MustPathParameters` - path parameters that must be present in the request
- `MustQueryParameters` - query parameters that must be present in the request
- `MustFormParameters` - form parameters that must be present in the request
- `MustHeaders` - headers that must be present in the request
- `MustBody` - body that must be present in the request

### Response Preparation
- `SetStatus` - HTTP status code to return
- `SetHeaders` - headers to return in the response
- `SetBody` - body to return in the response
- `SetFile` - file to return in the response

## Example Template

```json
{
    "Name": "testing",
    "Path": "/testing/:user_uuid",
    "Cases": {
        "GET": [
            {
                "MatchRequest": {
                    "MustPathParameters": {
                        "user_uuid": "${...}"
                    },
                    "MustHeaders": {},
                },
                "SetResponse": {
                    "SetStatus": 200,
                    "SetBody": {
                        "user_uuid": "${req.queryParams:user_uuid}",
                        "username": "x0rx3"
                    }
                }
            },
            {
                "MatchRequest": {
                    "MustQueryParameters": {
                        "download": "${...}"
                    },
                    "MustFormParameters": {},
                    "MustHeaders": {},
                    "MustBody": {}
                },
                "SetResponse": {
                    "SetStatus": 200,
                    "SetFile": "path_to_file/file"
                }
            },
            {
                "MatchRequest": {
                    "MustQueryParameters": {
                        "download": "${...}"
                    },
                },
                "SetResponse": {
                    "SetStatus": 200,
                    "SetFile": "testfile.txt"
                }
            },

            {
                "MatchRequest": {},
                "SetResponse": {
                    "SetStatus": 200,
                    "SetBody": {
                        "user_uuid": "${req.queryParams:user_uuid}",
                        "username": "x0rx3"
                    }
                }
            }
        ]
    }
}
```