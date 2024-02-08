run:
	air -c .air.toml

# build:
# 	@go build -o bin/eng-dev ./cmd/api/main.go

test:
	@go test -v -cover -short ./...
