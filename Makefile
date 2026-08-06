BINARY_NAME=baguette

GOOS ?= linux
GOARCH ?= amd64
EXE ?=

export GOOS
export GOARCH

.PHONY: build build\:win11 run test fmt clean install

build:
	go build -o bin/$(BINARY_NAME)$(EXE) .

build\:win11: GOOS=windows
build\:win11: GOARCH=386
build\:win11: EXE=.exe
build\:win11: build

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