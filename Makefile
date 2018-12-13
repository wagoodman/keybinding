all: clean build

build:
	go build

test: build
	go test -cover -v ./...

validate:
	@! gofmt -s -d -l . 2>&1 | grep -vE '^\.git/'
	go vet ./...

lint: build
	golint -set_exit_status $$(go list ./...)

clean:
	rm -rf build
	rm -rf vendor
	go clean

.PHONY: build install test lint clean release validate

