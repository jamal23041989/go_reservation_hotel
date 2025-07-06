build:
	@go build -o bin/api ./cmd

run: build
	@./bin/api


test:
	@go test -v ./...