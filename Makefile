VERSION = 0.0.1

.PHONY: all

all: clean build

build:
	go build -ldflags="-X main.version=${VERSION}"  -o build/ ./cmd/...

clean:
	rm -rf build