APP=scout-cli

.PHONY: build test run clean docker-build lint

build:
	go build -o bin/$(APP) ./cmd/scout-cli

test:
	go test ./...

run:
	go run ./cmd/scout-cli help

lint:
	go vet ./...

docker-build:
	docker build -t $(APP):latest .

clean:
	rm -rf bin