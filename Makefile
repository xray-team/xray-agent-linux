GO_BIN ?= go
ENV_BIN ?= env
OUT_BIN = xray-agent

export PATH := $(PATH):/usr/local/go/bin

all: clean build

download:
	$(ENV_BIN) GOPROXY=direct $(GO_BIN) get
	$(GO_BIN) mod tidy

update:
	$(ENV_BIN) GOPROXY=direct $(GO_BIN) get -u
	$(GO_BIN) mod tidy

test:
	$(GO_BIN) test -failfast ./...

lint:
    # Install:
    # https://golangci-lint.run/usage/install/#local-installation
    # binary will be $(go env GOPATH)/bin/golangci-lint
    # curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1
	golangci-lint run ./...

clean:
	$(GO_BIN) clean
	rm -f $(OUT_BIN)

build:
	$(GO_BIN) mod tidy
	$(GO_BIN) build -o $(OUT_BIN) -v
