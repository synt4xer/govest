# Define variables
BINARY_NAME=govest_app
BUILD_DIR=build
CMD_DIR=cmd

# Default target: build the binary
build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) ./$(CMD_DIR)

# Run the application
run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

# Run the application with Air for live reloading
dev:
	air

# Run tests
test:
	go test ./...

# Clean up the binary
clean:
	rm -rf $(BUILD_DIR)

# Install dependencies (optional, Go modules usually handle this)
deps:
	go mod tidy

# Lint the code (requires golangci-lint or another linter installed)
lint:
	golangci-lint run

# Catch all for undefined targets
.PHONY: build run dev test clean deps lint
