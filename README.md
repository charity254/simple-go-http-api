# Simple Go HTTP API

A basic HTTP API built with Go that demonstrates fundamental web server concepts including JSON responses, query parameters, request body parsing, HTTP method validation, and in-memory data storage with concurrent access.

## What the Service Does

This is a simple HTTP server with multiple endpoints that return JSON responses:

- **`/`** - Welcome message
- **`/health`** - Returns the service health status
- **`/status`** - Returns service status with uptime information
- **`/hello`** - Returns a personalized greeting from query parameters (GET)
- **`/greet`** - Returns a personalized greeting from JSON body (POST)
- **`/users`** - User management endpoints (POST to create, GET to list/retrieve)

## Features

- ✅ JSON request and response handling
- ✅ Query parameter parsing
- ✅ Request body parsing with validation
- ✅ HTTP method validation (GET/POST)
- ✅ Proper error handling with status codes
- ✅ Reusable error response function
- ✅ Service uptime tracking
- ✅ Thread-safe in-memory user storage with RWMutex**
- ✅ CRUD operations for user management**
- ✅ Clean code organization with separate handler file

## Project Structure
```
simple-go-http-api/
├── main.go       # Main server setup and routing
├── handlers.go   # HTTP handler functions and UserStore
├── test.http     # REST Client test file
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
   or 
```
   go run main.go handlers.go
```

3. You should see:
```
   Server starting on http://localhost:8080/
```

4. The server is now running at `http://localhost:8080`

## API Endpoints

### Root Endpoint

**Request:**
```bash
GET http://localhost:8080/
```

**Response:**
```json
{
  "message": "Welcome to the API!\n"
}
```

**Status Code:** 200 OK

---

### Health Check Endpoint

**Request:**
```bash
GET http://localhost:8080/health
```

**Response:**
```json
{
  "status": "healthy"
}
```

**Status Code:** 200 OK

---

### Status Endpoint (with Uptime)

**Request:**
```bash
GET http://localhost:8080/status
```

**Response:**
```json
{
  "service": "running",
  "uptime": "2h15m30.5s"
}
```

**Status Code:** 200 OK

The uptime shows how long the service has been running since it started.

---

### Hello Endpoint (Query Parameters - GET)

**Success Request:**
```bash
GET http://localhost:8080/hello?name=Alice
```

**Response:**
```json
{
  "message": "Hello, Alice!\n"
}
```

**Status Code:** 200 OK

**Error Request (missing name):**
```bash
GET http://localhost:8080/hello
```

**Response:**
```json
{
  "error": "Name is required"
}
```

**Status Code:** 400 Bad Request

---

### Greet Endpoint (JSON Body - POST)

**Success Request:**
```bash
POST http://localhost:8080/greet
Content-Type: application/json

{
  "name": "Alice"
}
```

**Response:**
```json
{
  "message": "Hello Alice"
}
```

**Status Code:** 200 OK

**Error Request (missing name):**
```bash
POST http://localhost:8080/greet
Content-Type: application/json

{
  "name": ""
}
```

**Response:**
```json
{
  "error": "Name is required"
}
```

**Status Code:** 400 Bad Request

**Error Request (invalid JSON):**
```bash
POST http://localhost:8080/greet
Content-Type: application/json

{invalid json}
```

**Response:**
```json
{
  "error": "Invalid JSON body"
}
```

**Status Code:** 400 Bad Request

**Error Request (wrong HTTP method):**
```bash
GET http://localhost:8080/greet
```

**Response:**
```json
{
  "error": "Only POST method is allowed"
}
```

**Status Code:** 405 Method Not Allowed

---

### User Management Endpoints

#### Create User (POST /users)

**Request:**
```bash
POST http://localhost:8080/users
Content-Type: application/json

{
  "name": "John Doe"
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "name": "John Doe",
  "created_at": "2026-02-20T10:30:00Z"
}
```

**Error Responses:**
- `400 Bad Request` - Invalid JSON or empty name
- `405 Method Not Allowed` - Wrong HTTP method

---

#### List All Users (GET /users)

**Request:**
```bash
GET http://localhost:8080/users
```

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "name": "John Doe",
    "created_at": "2026-02-20T10:30:00Z"
  },
  {
    "id": 2,
    "name": "Jane Smith",
    "created_at": "2026-02-20T10:31:00Z"
  }
]
```

Returns empty array `[]` if no users exist.

---

#### Get User by ID (GET /users?id={id})

**Request:**
```bash
GET http://localhost:8080/users?id=1
```

**Response (200 OK):**
```json
{
  "id": 1,
  "name": "John Doe",
  "created_at": "2026-02-20T10:30:00Z"
}
```

**Error Responses:**
- `400 Bad Request` - Invalid or missing ID
- `404 Not Found` - User doesn't exist

---

## Testing the API

### Option 1: Using the Browser (GET requests only)

- Root: `http://localhost:8080/`
- Health check: `http://localhost:8080/health`
- Status with uptime: `http://localhost:8080/status`
- Hello with name: `http://localhost:8080/hello?name=YourName`
- List all users: `http://localhost:8080/users`
- Get user by ID: `http://localhost:8080/users?id=1`

### Option 2: Using curl (Command Line)

**GET requests:**
```bash
curl http://localhost:8080/hello?name=Alice
curl http://localhost:8080/users
curl http://localhost:8080/users?id=1
```

**POST requests:**
```bash
curl -X POST http://localhost:8080/greet \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice"}'

curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe"}'
```

### Option 3: Using VS Code REST Client Extension

1. Install the "REST Client" extension in VS Code
2. Open the `test.http` file
3. Click "Send Request" above any request
4. View the response in the split panel

The `test.http` file includes test cases for all endpoints with various scenarios.

### Option 4: Using Postman

1. Download Postman from https://www.postman.com/downloads/
2. Create requests for each endpoint
3. Set the appropriate HTTP method (GET or POST)
4. Add request body for POST requests

## Response Format

All endpoints return JSON responses with appropriate headers:
- `Content-Type: application/json`

## Error Handling

The API uses standard HTTP status codes:

| Status Code | Meaning | When Used |
|-------------|---------|-----------|
| 200 | OK | Request succeeded |
| 201 | Created | Resource created successfully |
| 400 | Bad Request | Invalid input (missing parameters, invalid JSON) |
| 404 | Not Found | Resource not found |
| 405 | Method Not Allowed | Wrong HTTP method used |
| 500 | Internal Server Error | Server error (not currently implemented) |

All errors return a consistent JSON format:
```json
{
  "error": "Error message here"
}
```

## Code Structure

### Main Components

**`main.go`**
- Server initialization
- Route registration
- Port configuration

**`handlers.go`**
- Request handler functions
- JSON encoding/decoding
- Error handling logic
- UserStore implementation with thread-safe operations
- Data structures (User, UserStore, CreateUserRequest, StatusResponse, errorResponse)

**`test.http`**
- Test cases for all endpoints
- Example requests for development

### Key Functions

- `writeError()` - Reusable function for sending error responses
- `getRoot()` - Handles root endpoint
- `getHealth()` - Health check endpoint
- `getStatus()` - Service status with uptime
- `getHello()` - Greeting from query parameters
- `postGreet()` - Greeting from JSON body with POST validation
- `handleUsers()` - Dispatcher for user management endpoints
- `createUser()` - Creates a new user
- `listUsers()` - Returns all users
- `getUser()` - Returns a single user by ID

### UserStore

Thread-safe in-memory storage using `sync.RWMutex`:
- `Create(name string) *User` - Creates and stores a new user
- `GetById(id int64) (*User, bool)` - Retrieves a user by ID
- `List() []*User` - Returns all users

## Stopping the Server

Press `Ctrl + C` in the terminal where the server is running.

## API Summary Table

| Endpoint | Method | Parameters | Content-Type | Success Response | Error Codes |
|----------|--------|------------|--------------|------------------|-------------|
| `/` | GET | None | application/json | `{"message": "Welcome..."}` | - |
| `/health` | GET | None | application/json | `{"status": "healthy"}` | - |
| `/status` | GET | None | application/json | `{"service": "running", "uptime": "..."}` | - |
| `/hello` | GET | `name` (query) | application/json | `{"message": "Hello, {name}!"}` | 400 |
| `/greet` | POST | `{"name": "..."}` (body) | application/json | `{"message": "Hello {name}"}` | 400, 405 |
| `/users` | POST | `{"name": "..."}` (body) | application/json | User object with ID | 400, 405 |
| `/users` | GET | None | application/json | Array of users | 405 |
| `/users` | GET | `id` (query) | application/json | Single user object | 400, 404, 405 |

## Learning Concepts Demonstrated

This project demonstrates:
- Setting up an HTTP server in Go
- Creating RESTful API endpoints
- Handling different HTTP methods (GET, POST)
- Parsing query parameters
- Parsing JSON request bodies
- Encoding JSON responses
- Error handling with proper status codes
- Code organization with multiple files
- Using structs for structured data
- Request validation
- Testing APIs with REST Client
- Thread-safe concurrent data access with sync.RWMutex**
- In-memory data persistence**
- CRUD operations**
- Method routing and dispatching**

