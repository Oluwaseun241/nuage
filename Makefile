# Build the application
run: build
	@./bin/nuage

build:
	@echo "Building..."
	
	@go build -o bin/nuage cmd/api/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./internal/repositories -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f ./bin/nuage

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: all build run test clean
