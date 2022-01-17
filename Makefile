.PHONY: amm all clean fmt lint

GOBIN = ./build/bin
GOCMD = env GO111MODULE=on GOPROXY=https://goproxy.io go

amm:
	$(GOCMD) run build/ci.go install ./cmd/amm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/amm\" to launch amm."

all:
	$(GOCMD) build -v ./...
	$(GOCMD) run build/ci.go install ./cmd/...
	@echo "Done building."
	@echo "Find binaries in \"$(GOBIN)\" directory."
	@echo ""
	@echo "Copy config.toml to \"$(GOBIN)\"."
	@cp params/config-*.toml $(GOBIN)

clean:
	$(GOCMD) clean -cache
	rm -fr $(GOBIN)/*

fmt:
	./gofmt.sh

lint:
	golangci-lint run ./...
