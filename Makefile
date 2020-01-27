default: build

.PHONY: clean
clean:
	rm -f _output/*

.PHONY: build
build: clean
	go build -o _output -mod vendor ./cmd/...
	tar -cvzf _output/vendor.tar.gz vendor
