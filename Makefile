PROJECT_NAME=fbdl

default: build

help:
	@echo "Build targets:"
	@echo "  all      Run fmt vet build."
	@echo "  build    Build binary."
	@echo "  default  Run build."
	@echo "Quality targets:"
	@echo "  fmt   Format files with go fmt."
	@echo "  lint  Lint files with golangci-lint."
	@echo "Test targets:"
	@echo "  test-all                Run all tests."
	@echo "  test                    Run go test."
	@echo "  test-instantiating      Run instantiating tests."
	@echo "  test-parsing            Run parsing tests."
	@echo "  test-registerification  Run registerification tests."
	@echo "Other targets:"
	@echo "  help                Print help message."
	@echo "  update-tree-sitter  Update tree-sitter source files."

all: lint fmt build

build:
	go build -v -o $(PROJECT_NAME) ./cmd/fbdl


fmt:
	go fmt ./...

lint:
	golangci-lint run


test:
	go test ./...

test-parsing:
	@./scripts/test-parsing.sh

test-instantiating:
	@./scripts/test-instantiating.sh

test-registerification:
	@./scripts/reg-tests.sh

test-all: test test-parsing test-instantiating test-registerification


install:
	cp $(PROJECT_NAME) /usr/bin

uninstall:
	rm /usr/bin/$(PROJECT_NAME)

update-tree-sitter:
	@./scripts/update-tree-sitter.sh
