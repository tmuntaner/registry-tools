NAME    = registry-tools
VERSION = 0.0.1

default: build

clean:
	rm -f _output/*

build: clean
	go build -o _output ./cmd/...
