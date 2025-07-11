build:
	@go build -o bin/api ./cmd

run: build
	@./bin/api

stop:
	@taskkill /F /IM api.exe || true

seed:
	@go run scripts/seed.go

test:
	@go test -v ./...