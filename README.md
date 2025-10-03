# api-test

Lightweight example Go HTTP server used for demonstrating a small API with login/logout endpoints and static file serving.

This project is a minimal demo showing how to:

- Serve static files from the project directory
- Expose simple JSON API endpoints for login and logout
- Use a small internal toolkit package for JSON serialization and helpers

## Project layout

- `main.go` - the HTTP server and route handlers
- `index.html` - an example static file served at `/`
- `go.mod`, `go.sum` - Go module files

## Requirements

- Go 1.18+ (module-aware Go)
- Network port 8080 available

## Dependency

This project depends on `github.com/Go-Golang-Training/toolkit/v2`. Make sure your module proxy can fetch it, or replace the import with your local toolkit copy if needed.

## Build and run

From the project root run:

```bash
go mod tidy
go run ./
```

The server will start on port `:8080` and serve static files from the project directory. You should see a log line like:

```
Starting server (Application) on :8080
```

## Endpoints

- `GET /` - serves static files (e.g., `index.html`)
- `POST /api/login` - accepts JSON payload with `username` and `password`
- `POST /api/logout` - returns a JSON logout confirmation

### /api/login

Request (JSON):

```json
{
  "username": "me@here.com",
  "password": "verysecret"
}
```

Successful response (HTTP 202 Accepted):

```json
{
  "error": false,
  "message": "You have been logged in"
}
```

Failed response (HTTP 401 Unauthorized):

```json
{
  "error": true,
  "message": "Invalid credentials"
}
```

Notes:
- The example credentials are hard-coded in `main.go` for demonstration: username `me@here.com` and password `verysecret`.

### /api/logout

Request: `POST /api/logout` (no body required)

Response (HTTP 202 Accepted):

```json
{
  "message": "You have been logged out"
}
```

## Examples (curl)

Login (successful):

```bash
curl -s -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"username":"me@here.com","password":"verysecret"}'
```

Logout:

```bash
curl -s -X POST http://localhost:8080/api/logout
```

## Notes and next steps

- This repository is a teaching/demo project. It intentionally keeps logic simple and contains hard-coded credentials. Do not use as-is in production.
- Suggested improvements: add environment-based configuration, proper authentication/storage for sessions, input validation, and tests.

If you'd like, I can also add a small Makefile, a Dockerfile, or create unit tests for the handlers.
# api-test
