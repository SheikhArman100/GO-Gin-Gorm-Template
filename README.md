# GoLang API Template

A robust and scalable REST API template built with Go, featuring user authentication, database integration, and best practices for production deployment.

## Project Structure

```
.
├── cmd/
│   └── api/                    # Application entry point
│       └── main.go             # Server initialization and graceful shutdown implementation
├── internal/                   # Private application code
│   ├── database/               # Database layer
│   │   ├── database.go         # MySQL connection and GORM configuration
│   │   └── database_test.go    # Database integration tests
│   ├── handler/                # HTTP request handlers
│   │   ├── auth.go             # Authentication handlers (signup, signin)
│   │   └── user.go             # User-related request handlers
│   ├── interface/              # Interface definitions for clean architecture
│   ├── model/                  # Data models
│   │   └── user.go             # User entity with GORM schema
│   ├── response/               # API response handling
│   │   ├── apiError.go         # Error response structures
│   │   └── sendResponse.go     # Response formatting utilities
│   └── server/                 # HTTP server setup
│       ├── routes.go           # API route definitions and middleware
│       ├── routes_test.go      # Route testing
│       └── server.go           # Server configuration and initialization
├── .air.toml                   # Air configuration for live reload
├── docker-compose.yml          # MySQL database container configuration
├── go.mod                      # Go module dependencies
└── Makefile                    # Development and build automation commands
```

## Features

- **Clean Architecture**: Organized with a clear separation of concerns
- **RESTful API**: Built using the Gin web framework
- **Database Integration**: MySQL support using GORM ORM
- **Authentication**: User signup and signin endpoints
- **CORS Support**: Configured for cross-origin requests
- **Graceful Shutdown**: Handles shutdown signals properly
- **Health Check**: Endpoint for monitoring service health
- **Docker Support**: Easy database setup with Docker
- **Hot Reload**: Development with automatic rebuilding

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.16 or higher
- Docker and Docker Compose
- Make

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```
Create DB container
```bash
make docker-run
```

Shutdown DB Container
```bash
make docker-down
```

DB Integrations Test:
```bash
make itest
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```
