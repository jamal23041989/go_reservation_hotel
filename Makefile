build:
	@go build -o bin/api ./cmd

run: build
	@./bin/api

stop:
	@taskkill /F /IM api.exe || true

seed:
	@go run scripts/seed.go

docker:
	echo "building docker file"
	@docker build -t api .
	echo "running API inside Docker container"
	@docker run -p 3000:3000 api

test:
	@go test -v ./...