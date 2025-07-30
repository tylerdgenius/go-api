BINARY_NAME := api-template

.PHONY: build
build:
	go build -o bin/$(BINARY_NAME) cmd/server/main.go

.PHONY: run
run: build
	./bin/$(BINARY_NAME)

.PHONY: fmt
fmt: 
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: clean
clean:
	rm -rf bin/$(BINARY_NAME)

.PHONY: tidy
tidy:
	go mod tidy