# HelloWorld Service

A simple Go HTTP service that demonstrates the CI/CD pipeline and code conventions.

## Features

- **RESTful API**: Simple HTTP endpoints for greeting functionality
- **Health Checks**: Built-in health check endpoint for monitoring
- **Metrics**: Basic metrics endpoint for observability
- **Configuration**: YAML-based configuration with environment variable overrides
- **Docker Support**: Containerized for easy deployment
- **Testing**: Comprehensive unit tests with coverage reporting

## API Endpoints

### Health Check

```http
GET /health
```

Returns the service health status and uptime information.

**Response:**

```json
{
  "status": "healthy",
  "uptime": "1h30m45s",
  "timestamp": "2024-01-15T10:30:00Z"
}
```

### Basic Hello

```http
GET /hello
```

Returns a basic "Hello, World!" greeting.

**Response:**

```json
{
  "message": "Hello, World!",
  "timestamp": "2024-01-15T10:30:00Z",
  "version": "1.0.0"
}
```

### Personalized Hello

```http
GET /hello/{name}
```

Returns a personalized greeting for the specified name.

**Response:**

```json
{
  "message": "Hello, Alice!",
  "timestamp": "2024-01-15T10:30:00Z",
  "version": "1.0.0"
}
```

### Metrics

```http
GET /metrics
```

Returns basic service metrics.

**Response:**

```json
{
  "requests": 42,
  "uptime": "1h30m45s",
  "status": "running"
}
```

## Configuration

The service can be configured via `config/config.yaml` or environment variables:

| Environment Variable | Description | Default |
|---------------------|-------------|---------|
| `PORT` | HTTP server port | `8080` |
| `READ_TIMEOUT` | HTTP read timeout (seconds) | `30` |
| `WRITE_TIMEOUT` | HTTP write timeout (seconds) | `30` |
| `IDLE_TIMEOUT` | HTTP idle timeout (seconds) | `120` |
| `LOG_LEVEL` | Logging level | `info` |
| `LOG_FORMAT` | Log format (json/text) | `json` |

## Development

### Prerequisites

- Go 1.21 or later
- Docker (for containerization)

### Running Locally

```bash
# Install dependencies
go mod download

# Run the service
go run cmd/main.go
```

The service will start on `http://localhost:8080`.

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -v -race -coverprofile=coverage.out ./...

# View coverage report
go tool cover -html=coverage.out
```

### Building

```bash
# Build binary
go build -o helloworld cmd/main.go

# Build Docker image
docker build -t helloworld .

# Run Docker container
docker run -p 8080:8080 helloworld
```

## Docker

The service includes a multi-stage Dockerfile that:

- Uses Go 1.21 Alpine for building
- Creates a minimal Alpine runtime image
- Runs as a non-root user for security
- Includes health checks
- Exposes port 8080

### Docker Commands

```bash
# Build image
docker build -t helloworld .

# Run container
docker run -p 8080:8080 helloworld

# Run with custom port
docker run -p 9090:8080 -e PORT=8080 helloworld
```

## CI/CD Integration

This service is designed to work with the GitHub Actions CI/CD pipeline defined in the repository. The pipeline will:

1. **Detect Changes**: Monitor changes to the `./src` directory
2. **Run Tests**: Execute unit tests and generate coverage reports
3. **Code Quality**: Run static analysis, formatting checks, and security scans
4. **Build Container**: Create and scan Docker images
5. **Deploy**: Push images to GitHub Container Registry

The service follows conventional commit standards for automated versioning and release management.

## Project Structure

```text
src/
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── config/
│   │   ├── config.go        # Configuration management
│   │   └── config_test.go   # Configuration tests
│   └── handlers/
│       ├── handlers.go      # HTTP handlers
│       └── handlers_test.go # Handler tests
├── config/
│   └── config.yaml          # Default configuration
├── Dockerfile               # Container definition
├── go.mod                   # Go module definition
├── go.sum                   # Dependency checksums
└── README.md               # This file
```

## Contributing

1. Follow conventional commit message format
2. Write tests for new functionality
3. Ensure all tests pass and code is formatted
4. Update documentation as needed

The CI/CD pipeline will validate all changes and provide feedback on pull requests.
