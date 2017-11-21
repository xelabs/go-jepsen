export GOPATH := $(shell pwd)
export PATH := $(GOPATH)/bin:$(PATH)

build:
	@echo "--> go get..."
	go get github.com/XeLabs/go-mysqlstack/driver

	@echo "--> Building..."
	@mkdir -p bin/
	go build -v -o bin/jepsen src/cli/jepsen.go
	@chmod 755 bin/*

clean:
	@echo "--> Cleaning..."
	@go clean
	@rm -f bin/*

fmt:
	go fmt ./...

test:
	@echo "--> Testing..."
	@$(MAKE) testxworker
	@$(MAKE) testxcommon
	@$(MAKE) testsnapshot

testxworker:
	go test -v xworker
testxcommon:
	go test -v xworker
testsnapshot:
	go test -v snapshot

# code coverage
COVPKGS =	xcmd\
			xworker\
			xcommon\
			snapshot
coverage:
	go build -v -o bin/gotestcover \
	src/vendor/github.com/pierrre/gotestcover/*.go;
	gotestcover -coverprofile=coverage.out -v $(COVPKGS)
	go tool cover -html=coverage.out
.PHONY: build clean fmt test coverage
