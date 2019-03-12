GO_BIN ?= go

install:
	@packr2 clean
	@packr2
	@$(GO_BIN) install -v ./.
	@make tidy

clean:
	@rm -rf dist/
	@rm -rf ./devlx

reset: install
	@rm -rf ~/.devlx
	@devlx config -t
	@devlx profile gui -w
	@devlx profile cli -w

tidy:
ifeq ($(GO111MODULE),on)
	@$(GO_BIN) mod tidy
endif

deps:
	@$(GO_BIN) get github.com/gobuffalo/release
	@$(GO_BIN) get github.com/gobuffalo/packr/v2/packr2
	@$(GO_BIN) get github.com/gobuffalo/shoulders
	@$(GO_BIN) get ./...
	@make tidy

build:
	@packr2
	@$(GO_BIN) build -v .
	@make tidy

test:
	@packr2
	@$(GO_BIN) test ./...
	@make tidy

shoulders:
	@shoulders -n github.com/bketelsen/devlx -w

ci-deps:
	$(GO_BIN) get -t ./...

ci-test:
	$(GO_BIN) test -race ./...

lint:
	@gometalinter --vendor ./... --deadline=1m --skip=internal
	@make tidy

update:
	@$(GO_BIN) get -u
	@make tidy
	@packr2
	@make test
	@make install
	@make tidy

release-test:
	@$(GO_BIN) test -race ./...
	@make tidy

release:
	@make tidy
	@make shoulders
	@release -y -f ./cmd/version.go
	@make tidy
