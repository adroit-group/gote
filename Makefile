COMMITTISH := $(shell git rev-parse --short HEAD)
BUILD_DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

export COMMITTISH
export BUILD_DATE

default: help

# Prints the available commands and their descriptions
help:
	@echo "Available commands:"
	@awk 'BEGIN {FS = ":.*"; printf "%-20s %s\n", "Target", "Description"; printf "%-20s %s\n", "------", "-----------"} \
		/^[#]+ .*$$/ {desc = (desc ? desc " " : "") substr($$0, 3)} \
		/^[a-zA-Z0-9_-]+:/ {printf "%-20s %s\n", $$1, desc; desc = ""}' $(MAKEFILE_LIST)

# Starts the project
start:
	@docker-compose up -d --build

# Stops the project and removes the containers
clean:
	@docker-compose down --volumes --remove-orphans

# Runs the tests for the project
test:
	@go test -cover ./...

# Lints the code using golangci-lint and `go fmt`, and fix the `go.mod` file with `go mod tidy`
lint:
	@go mod tidy
	@go fmt ./...
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint run ./...

