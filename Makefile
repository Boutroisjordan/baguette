BINARY_NAME=baguette

.PHONY: build, run, test, fmt ,clean, install
build:
	go build -o bin/$(BINARY_NAME) .

run:
	go run .

test:
	go test ./...

fmt:
	go fmt ./...

clean:
	rm -rf bin/

install:
	go install .