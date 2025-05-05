# Go CRUD API with Gin

A simple RESTful API demonstrating CRUD operations using the Gin framework in Go.

## Project Structure

```
.
├── cmd/             # Application entry points
│   └── api/         # API server entry point
│       └── main.go  # Main application file
├── internal/        # Private application code
│   ├── handler/     # HTTP handlers
│   ├── middleware/  # HTTP middleware
│   ├── model/       # Data models
│   ├── repository/  # Data access layer
│   └── service/     # Business logic
├── pkg/             # Public libraries
│   └── utils/       # Utility functions
├── config/          # Configuration files
├── docs/            # Documentation
│   └── swagger/     # API documentation
├── .env.example     # Example environment variables
├── .gitignore       # Git ignore file
├── go.mod           # Go module file
├── go.sum           # Go module checksum file
└── README.md        # This file
```

## Requirements

- Go 1.16+
- Docker (optional, for containerization)

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/yourusername/go-crud-api.git
cd go-crud-api
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables:
```bash
cp .env.example .env
# Edit .env file with your configuration
```

4. Run the application:
```bash
go run cmd/api/main.go
```

The server will start on `http://localhost:8080`.

## API Endpoints

| Method | URL | Description |
|--------|-----|-------------|
| GET    | /api/v1/users | Get all users |
| GET    | /api/v1/users/:id | Get a user by ID |
| POST   | /api/v1/users | Create a new user |
| PUT    | /api/v1/users/:id | Update a user |
| DELETE | /api/v1/users/:id | Delete a user |

## Using the API Endpoints

### Get Health Status
```bash
curl -X GET http://localhost:8080/health
```

### Create a User
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@example.com"
  }'
```

### Get All Users
```bash
curl -X GET http://localhost:8080/api/v1/users
```

### Get User by ID
```bash
# Replace USER_ID with the actual user ID
curl -X GET http://localhost:8080/api/v1/users/USER_ID
```

### Update a User
```bash
# Replace USER_ID with the actual user ID
curl -X PUT http://localhost:8080/api/v1/users/USER_ID \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Johnny",
    "last_name": "Doe",
    "email": "johnny.doe@example.com"
  }'
```

### Delete a User
```bash
# Replace USER_ID with the actual user ID
curl -X DELETE http://localhost:8080/api/v1/users/USER_ID
```

## Testing

Run the tests with:
```bash
go test ./...
```

## Docker Support

Build the Docker image:
```bash
docker build -t go-crud-api .
```

Run the container:
```bash
docker run -p 8080:8080 go-crud-api
``` 