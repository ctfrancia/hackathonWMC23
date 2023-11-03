.PHONY: help build clean run test coverage lint fmt vet install uninstall

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  help        Show this help message"
	@echo "  build       Build the project"
	@echo "  clean       Clean the project"
	@echo "  run         Run the project"
	@echo "  test        Run the tests"
	@echo "  coverage    Run the tests with coverage"
	@echo "  lint        Run the linter"
	@echo "  fmt         Run the formatter"
	@echo "  vet         Run the vet"
	@echo "  install     Install the project"
	@echo "  uninstall   Uninstall the project"

build:
	@echo "Building the project..."
	@go build -o bin/$(PROJECT) $(PROJECT).go

clean:
	@echo "Cleaning the project..."
	@rm -rf bin/$(PROJECT)

run:
	@echo "Running the project..."
	@go run $(PROJECT).go

test:
	@echo "Running the tests..."
	@go test -v ./...

coverage:
	@echo "Running the tests with coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

lint:
	@echo "Running the linter..."
	@golangci-lint run

fmt:
	@echo "Running the formatter..."
	@gofmt -s -w .

vet:
	@echo "Running the vet..."
	@go vet ./...

install:
	@echo "Installing the project..."
	@go install

uninstall:
	@echo "Uninstalling the project..."
	@go clean -i
