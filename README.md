# GoTe, a Go Microservice Template

[![Go Version](https://img.shields.io/badge/Go-1.24%2B-blue.svg)](https://golang.org/doc/devel/release.html)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

A streamlined Go microservice template for quickly bootstrapping new microservices with consistent structure and tooling. This template helps eliminate repetitive boilerplate code and maintains standardized practices across multiple services.

## Features

- **Minimal Boilerplate**: Focus on writing business logic instead of infrastructure code
- **Clean Architecture**: Organized folder structure following Go best practices
- **Built-in Observability**: Pre-configured OpenTelemetry integration for tracing and metrics, structured logging with `slog`.
- **Standard Health Endpoints**: Ready-to-use health check and version endpoints
- **Secure by Default**: Builds onto Chainguard static base image for minimal attack surface
- **Developer-Friendly**: Complete with Task-based workflows and Docker Compose setup

## Project Structure

```txt
.
├── api/            # OpenAPI 3.0/3.1 documentation
├── build/          # Build files and Dockerfile
├── cmd/            # Application entry points
├── configs/        # Configuration files
├── internal/       # Private application code
│   └── httpserver/ # HTTP server implementation
├── pkg/            # Public libraries (imported from upstream after templating)
├── hack/           # Scripts for development and template cleanup
└── .github/        # GitHub Actions and PR templates
```

## Requirements

- Go 1.23+ (1.24 recommended)
- [Task](https://taskfile.dev/)
- [Mise](https://github.com/jdx/mise) (optional)
- Docker and Docker Compose

## Getting Started

### Using This Template

The template can be used to create a new microservice with the following steps:

```bash
# The setup script will:
# 1. Clone this repository
# 2. Delete the pkg directory and git history
# 3. Initialize a new git repository
# 4. Replace package names throughout the codebase
# 5. Create an initial commit

./hack/setup-new-project.sh
```

### Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/go-microservice-template.git
cd go-microservice-template

# Install dependencies and tools
task get-tools

# Or use Mise
mise install
```

### Configuration

This project uses [Viper](https://github.com/spf13/viper) for configuration. Configuration files are located in the `configs/` directory.
You can create new configuration options by extending the `internal/config.go` file.
Environment variables can be used to override configuration values.

#### Environment Variables

The template includes OpenTelemetry instrumentation with the following environment variables:

```env
OTEL_SERVICE_NAME: yourapp

# OTLP collector endpoint
OTEL_EXPORTER_OTLP_ENDPOINT: http://otel-collector:4317

# available options: none, console, zipkin, otlp
OTEL_TRACES_EXPORTER: console

# available options: none, console, prometheus, otlp
OTEL_METRICS_EXPORTER: prometheus

# only used in otlp mode, available options: http, grpc
OTEL_EXPORTER_OTLP_TRACES_PROTOCOL: grpc
```

Additional configuration parameters can be found in `internal/config.go`.

### Running Locally

```bash
# Start development environment with Docker Compose
task dev

# Or run directly
task run
```

### Testing

```bash
task test
```

## API Endpoints

The template includes the following built-in endpoints:

- `/__health__` - Health check endpoint for infrastructure monitoring
- `/__version__` - Returns version information about the running service

All API endpoints should be documented in the OpenAPI specifications in the `api/` directory.

## Tech Stack

- [go-chi](https://github.com/go-chi/chi) - HTTP routing
- [slog](https://pkg.go.dev/log/slog) - Structured logging
- [go-playground/validator](https://github.com/go-playground/validator) - Request validation
- [testify](https://github.com/stretchr/testify) - Testing framework
- [opentelemetry-go-auto-instrumentation](https://github.com/alibaba/opentelemetry-go-auto-instrumentation)
  - Telemetry
  - You can call `traceId, spanId := trace.GetTraceAndSpanId()` from `"go.opentelemetry.io/otel/sdk/trace"` anywhere to obtain the trace and span IDs.
  - The full instrumentation happens compile-time, so you don't have to worry about writing any more instrumentation code.
- [viper](https://github.com/spf13/viper) - Configuration management
- [Task](https://taskfile.dev/) - Task automation

## Security

This project builds into a minimal [Chainguard](https://www.chainguard.dev/) static base image for security and reduced attack surface.

## Implementation Examples

For examples of how to use the HTTP server implementation, refer to the `internal/httpserver/server.go` file.

## Future Enhancements

- Terraform deployment configurations
- Additional security integrations
- Postgres database integration
- Redis caching integration
- Expanded service configurations in Docker Compose
- Trivy security scanning
- SALSA integration for software supply chain security

## Contributing

Contribution guidelines are coming soon!

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.
