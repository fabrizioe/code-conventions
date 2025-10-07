# Copilot Instructions for Industrial IoT System

## Project Overview
This is an Industrial IoT (IIoT) platform for device management and monitoring built with microservices architecture. The system handles real-time telemetry data from edge devices through a message broker to time-series storage and analytics.

## Architecture
The system follows a hub-and-spoke pattern with these key components:
- **Edge Gateway**: Collects data from field devices and forwards to protocol adapters
- **Field Protocol Adapters**: Handle heterogeneous protocols (Modbus, OPC-UA, MQTT, HTTP) and normalize data
- **AuthN/Z**: JWT-based authentication via Keycloak OIDC provider
- **Data Ingestion**: Validates, normalizes and routes device data to components
- **Device Manager**: Distributed system for centralized device configuration, orchestration and command execution
- **Device Monitoring**: Real-time monitoring, alerting, and performance tracking
- **Time-Series DB**: Stores telemetry with timestamp-based queries
- **Analytics**: ML-based predictive maintenance and reporting

## Development Conventions

### Project Structure
- Each service has its own `Spec.md` with domain entities and responsibilities
- Implementation code lives in separate directories from specifications
- Docker Compose used for local development stack
- Go services follow standard project layout: `cmd/`, `internal/`, `config/`

### Git Workflow
- Use trunk-based development with feature branches
- Branch naming: `<task-id>-<short-description>`
- Conventional commit messages required
- All tests must pass before merge to main

### Configuration Management
- YAML-based configuration files in `config/` directories
- Environment-specific overrides supported
- Structured logging with configurable levels (debug/info/error)
- JSON formatter for production environments

### Authentication & Authorization
- JWT tokens for device authentication via Keycloak
- OIDC flow for user authentication
- Device provisioning service manages short-lived JWT tokens
- Eclipse Hono handles device registry and ACL

## Key Implementation Patterns

### Service Structure (Go)
```go
// Standard service initialization pattern
keycloakService := services.NewKeycloakService(&cfg.Keycloak, logger)
honoService := services.NewHonoService(&cfg.Hono, logger)
provisioningService := services.NewProvisioningService(cfg, keycloakService, honoService, logger)
```

### Data Flow Events
- Use standardized JSON event format with `deviceId`, `timestamp`, `measurements`
- Event types: `telemetry_data`, `device_command`, `command_response`, `threshold_exceeded`
- Quality indicators included in telemetry events (`good`, `bad`, `uncertain`)

### Messaging Patterns
- Protocol adaptation for heterogeneous field devices (Modbus, OPC-UA, MQTT, HTTP)
- Data normalization and JSON standardization by Field Protocol Adapters
- Event-driven alerting with configurable thresholds via Device Monitoring
- Command/response for device control operations via Device Manager
- Real-time data distribution from Data Ingestion to Device Manager and Device Monitoring

## Development Workflow

### Local Development
1. Start stack: `docker-compose up -d` in `AuthNZ/` directory
2. Setup Keycloak: `./scripts/setup-keycloak.sh`
3. Setup Hono: `./scripts/setup-hono.sh`
4. Test connectivity: `./scripts/test-mqtt.sh` or `./scripts/test-http.sh`

### Service Endpoints
- Keycloak Admin: http://localhost:8080 (admin/admin)
- Hono Device Registry: http://localhost:28080
- Hono MQTT: localhost:1883
- Hono HTTP: http://localhost:8088
- Kafka UI: http://localhost:8089

### Testing
- Use provided test scripts for MQTT and HTTP connectivity
- Validate JWT token flow through device provisioning service
- Test device registration and telemetry ingestion end-to-end

## Critical Integration Points

### Device Onboarding
1. Device registers with Device Manager provisioning service
2. Keycloak issues JWT token with device claims
3. Device Manager Agent (peripheral component) manages device state machine (NEW→RUNNING→MAINTENANCE)
4. Field Protocol Adapters handle protocol-specific communication
5. Device can send telemetry and receive commands through appropriate adapters

### Data Pipeline
1. Edge Gateway → Field Protocol Adapters (protocol conversion)
2. Field Protocol Adapters → Data Ingestion Service (normalization)
3. Data Ingestion Service → Time-Series Database (historical storage)
4. Data Ingestion Service → Device Manager (command orchestration)
5. Data Ingestion Service → Device Monitoring (real-time monitoring)
6. Analytics Service queries historical data from Time-Series Database

### Error Handling
- Structured logging with error context
- Graceful degradation for service dependencies
- Backpressure handling in data ingestion
- Circuit breaker patterns for external service calls

When implementing new features, maintain consistency with existing service patterns, use the established configuration approach, and ensure proper integration with the OIDC authentication flow.