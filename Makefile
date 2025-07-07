build:
	@go build -o bin/api ./cmd

run: build
	@./bin/api

seed:
	@go run scripts/seed.go

test:
	@go test -v ./...