.PHONY: run build test docker-up docker-down clean deps

run:
	go run cmd/api/main.go

build:
	go build -o bin/api cmd/api/main.go

test:
	go test -v ./...

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

docker-rebuild:
	docker-compose up -d --build

deps:
	go mod download
	go mod tidy

clean:
	rm -rf bin/
	go clean

.DEFAULT_GOAL := run
