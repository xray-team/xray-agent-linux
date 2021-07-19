GO_BIN ?= go
ENV_BIN ?= env
OUT_BIN = xray-agent

export PATH := $(PATH):/usr/local/go/bin

download:
	$(ENV_BIN) GOPROXY=direct $(GO_BIN) get
	$(GO_BIN) get github.com/golangci/golangci-lint/cmd/golangci-lint
	$(GO_BIN) mod tidy

update:
	$(ENV_BIN) GOPROXY=direct $(GO_BIN) get -u
	$(GO_BIN) get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	$(GO_BIN) mod tidy

test:
	$(GO_BIN) test -failfast ./...

lint:
	golangci-lint run ./...

clean:
	$(GO_BIN) clean
	rm -f $(OUT_BIN)

build:
	$(GO_BIN) mod tidy
	$(GO_BIN) build -o $(OUT_BIN) -v
