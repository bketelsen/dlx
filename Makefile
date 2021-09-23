GO_BIN ?= go


.PHONY: install
install:
	@$(GO_BIN) install -v ./.

clean:
	@rm -rf dist/
	@rm -rf ./dlx

reset: install
	@rm -rf ~/.dlx
	@dlx config -t

tidy:
	@$(GO_BIN) mod tidy

deps:
	@$(GO_BIN) get ./...

build:
	@$(GO_BIN) build -v .

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

server-install:
	./scripts/lxd.sh
	./scripts/distrobuilder.sh
	./scripts/debootstrap.sh
	chmod +x ./scripts/devices && sudo cp ./scripts/devices /usr/local/bin/