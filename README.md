# Simple Go HTTP API

A basic HTTP API built with Go that demonstrates fundamental web server concepts.

## What the Service Does

This is a simple HTTP server with two endpoints:

- **`/health`** - Returns the service status
- **`/hello`** - Returns a personalized greeting based on a name parameter

## How to Run It

### Prerequisites
- Go installed on your system (version 1.16 or higher)

### Steps

1. Save the code in a file called `main.go`

2. Run the server:
   ```bash
   go run main.go
   ```

3. You should see:
   ```
  Server starting on http://localhost:8080/
   ```

4. The server is now running at `http://localhost:8080`

## Example Requests

### Health Check Endpoint

**Request:**
```bash
curl http://localhost:8080/health
```

**Response:**
```
Service is running
```

---

### Hello Endpoint (with name parameter)

**Request:**
```bash
curl "http://localhost:8080/hello?name=Alice"
```

**Response:**
```
Hello, Alice!
```

**Another example:**
```bash
curl "http://localhost:8080/hello?name=John"
```

**Response:**
```
Hello, John!
```

---

### Hello Endpoint (without name parameter)

**Request:**
```bash
curl http://localhost:8080/hello
```

**Response:**
```
Error: name parameter is required
```

**Status Code:** 400 Bad Request

---

## Testing in Browser

You can also test these endpoints directly in your web browser:

- Health check: `http://localhost:8080/health`
- Hello with name: `http://localhost:8080/hello?name=YourName`

## Stopping the Server

Press `Ctrl + C` in the terminal where the server is running.
