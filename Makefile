PROJECT_NAME=fbdl

default: build

all: fmt vet build

build:
	go build -v -o $(PROJECT_NAME) .

help:
	@echo "Build targets:"
	@echo "  all      Run fmt vet build."
	@echo "  build    Build binary."
	@echo "  default  Run build."
	@echo "Quality targets:"
	@echo "  fmt       Format files with go fmt."
	@echo "  vet       Examine go sources with go vet."
	@echo "  errcheck  Examine go sources with errcheck."
	@echo "Test targets:"
	@echo "  test-all            Run all tests."
	@echo "  test                Run go test."
	@echo "  test-instantiating  Run instantiating tests."
	@echo "  test-parsing        Run parsing tests."
	@echo "Other targets:"
	@echo "  help                Print help message."
	@echo "  update-tree-sitter  Update tree-sitter source files."

fmt:
	go fmt ./...

vet:
	go vet ./...

errcheck:
	errcheck -verbose ./...

test-instantiating:
	@./scripts/test-instantiating.sh

test-parsing:
	@./scripts/test-parsing.sh

test:
	go test ./...

test-all: test test-parsing test-instantiating

install:
	cp $(PROJECT_NAME) /usr/bin

uninstall:
	rm /usr/bin/$(PROJECT_NAME)

update-tree-sitter:
	@./scripts/update-tree-sitter.sh
