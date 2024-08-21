// Make sure to create an fresh_config 

start:
	@go run cmd/api/main.go

build:
	@echo "Building..."
	@go build -o main.exe cmd/api/main.go
dev:
	@echo "Running dev mode..."
	@fresh -c fresh_env.config
	
test:
	@echo "Running tests..."
	@go test ./... -v