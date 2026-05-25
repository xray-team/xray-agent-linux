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
    # https://golangci-lint.run/docs/welcome/install/local/
    # binary will be $(go env GOPATH)/bin/golangci-lint
    # curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b $(go env GOPATH)/bin latest
    # golangci-lint --version
	golangci-lint run ./...

clean:
	$(GO_BIN) clean
	rm -f $(OUT_BIN)

build:
	$(GO_BIN) mod tidy
    # Build a dynamic binary (with libc6/glibc dependency)
    # $(GO_BIN) build -o $(OUT_BIN) -v
    # Build a fully static binary (no libc6/glibc dependency)
	$(ENV_BIN) CGO_ENABLED=0 GOOS=linux $(GO_BIN) build -trimpath -ldflags=' -s -w -extldflags "-static"' -o $(OUT_BIN) -v
