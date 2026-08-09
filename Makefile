BINARY_NAME=baguette

GOOS ?= linux
GOARCH ?= amd64
EXE ?=

export GOOS
export GOARCH

.PHONY: build build\:win11 build\:macos build\:darwin run test fmt clean install

build:
	echo '$(BINARY_NAME) $(EXE) $(GOOS) $(GOARCH)' && go build -o bin/$(BINARY_NAME)$(EXE) .

build\:win11: GOOS=windows
build\:win11: GOARCH=386
build\:win11: EXE=.exe
build\:win11: build

build\:macos: GOOS=linux
build\:macos: GOARCH=arm
build\:macos: build

# Works for Macos on ARM
build\:darwin: GOOS=darwin
build\:darwin: GOARCH=amd64
build\:darwin: build

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