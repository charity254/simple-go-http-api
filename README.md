# Simple Go HTTP API

A basic HTTP API built with Go that demonstrates fundamental web server concepts with JSON responses.

## What the Service Does

This is a simple HTTP server with multiple endpoints that return JSON responses:

- **`/`** - Welcome message
- **`/health`** - Returns the service health status
- **`/status`** - Returns service status with uptime information
- **`/hello`** - Returns a personalized greeting based on a name parameter

## Project Structure

```
project/
├── main.go       # Main server setup and routing
├── handlers.go   # HTTP handler functions
└── README.md     # This file
```

## How to Run It

### Prerequisites
- Go installed on your system (version 1.16 or higher)

### Steps

1. Make sure you have both `main.go` and `handlers.go` in the same directory

2. Run the server:
   ```bash
   go run .
   ```

3. You should see:
   ```
   Server starting on http://localhost:8080/
   ```

4. The server is now running at `http://localhost:8080`

## Example Requests

### Root Endpoint

**Request:**
```bash
curl http://localhost:8080/
```

**Response:**
```json
{
  "message": "Welcome to the API!\n"
}
```

---

### Health Check Endpoint

**Request:**
```bash
curl http://localhost:8080/health
```

**Response:**
```json
{
  "status": "healthy"
}
```

---

### Status Endpoint (with Uptime)

**Request:**
```bash
curl http://localhost:8080/status
```

**Response:**
```json
{
  "service": "running",
  "uptime": "2h15m30.5s"
}
```

The uptime shows how long the service has been running since it started.

---

### Hello Endpoint (with name parameter)

**Request:**
```bash
curl "http://localhost:8080/hello?name=Alice"
```

**Response:**
```json
{
  "message": "Hello, Alice!\n"
}
```

**Another example:**
```bash
curl "http://localhost:8080/hello?name=John"
```

**Response:**
```json
{
  "message": "Hello, John!\n"
}
```

---

### Hello Endpoint (without name parameter)

**Request:**
```bash
curl http://localhost:8080/hello
```

**Response:**
```json
{
  "error": "name parameter is required"
}
```

**Status Code:** 400 Bad Request

---

## Testing in Browser

You can also test these endpoints directly in your web browser:

- Root: `http://localhost:8080/`
- Health check: `http://localhost:8080/health`
- Status with uptime: `http://localhost:8080/status`
- Hello with name: `http://localhost:8080/hello?name=YourName`

## Response Format

All endpoints return JSON responses with appropriate `Content-Type: application/json` headers.

## Stopping the Server

Press `Ctrl + C` in the terminal where the server is running.

## API Summary

| Endpoint | Method | Parameters | Success Response | Error Response |
|----------|--------|------------|------------------|----------------|
| `/` | GET | None | `{"message": "Welcome..."}` | - |
| `/health` | GET | None | `{"status": "healthy"}` | - |
| `/status` | GET | None | `{"service": "running", "uptime": "..."}` | - |
| `/hello` | GET | `name` (required) | `{"message": "Hello, {name}!"}` | `{"error": "name parameter is required"}` (400) |
