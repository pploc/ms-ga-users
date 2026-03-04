.PHONY: build run test clean docker-build docker-run tidy fmt lint db-up

APP_NAME=ms-ga-user
PORT=8083

build:
	go build -o bin/$(APP_NAME) cmd/api/main.go

run:
	go run cmd/api/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/

tidy:
	go mod tidy

docker-build:
	docker build -t $(APP_NAME) .

docker-run:
	docker run -p $(PORT):$(PORT) --env-file .env $(APP_NAME)

fmt:
	gofmt -w .
	goimports -w ./...

lint: fmt
	gofmt -l . | grep . && exit 1 || true
