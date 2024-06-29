default: build

.PHONY: build
build:
	@echo "Building binary..."
	@go build cmd/main.go

.PHONY: generate
generate:
	@echo "Generating code..."
	@go generate ./...

.PHONY: run
run:
	@echo "Running server..."
	@go run cmd/main.go

.PHONY: fmt
fmt:
	@echo "Formatting code..."
	@go fmt ./...

.PHONY: test
test:
	@echo "Running tests..."
	@go test ./... -v
