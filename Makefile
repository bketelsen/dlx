GO_BIN ?= go


.PHONY: install
install:
	@$(GO_BIN) install -v ./.

clean:
	@rm -rf dist/

reset: install
	@rm -rf ~/.dlx

tidy:
	@$(GO_BIN) mod tidy

deps:
	@$(GO_BIN) get ./...
	@$(GO_BIN) install github.com/gohugoio/hugo@latest
	@$(GO_BIN) install github.com/goreleaser/goreleaser@latest

build:
	@$(GO_BIN) build -v -o bin/dlx 

docs: build
	./bin/dlx docs

test:
	@$(GO_BIN) test ./...

ci-deps:
	$(GO_BIN) get -t ./...

ci-test:
	$(GO_BIN) test -race ./...


update:
	@$(GO_BIN) get -u
	@make tidy
	@make test
	@make install

release-test:
	@$(GO_BIN) test -race ./...
