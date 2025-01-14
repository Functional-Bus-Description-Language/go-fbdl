# Project name
NAME=fbdl

default: build

help:
	@echo "Build targets:"
	@echo "  all      Run lint fmt build."
	@echo "  build    Build binary."
	@echo "  debug    Build binary for debugging."
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
	@echo "  test-expr               Run expression evaluation tests."
	@echo "Other targets:"
	@echo "  help                Print help message."


# Build targets
all: lint fmt build

build:
	go build -v -o $(NAME) ./cmd/$(NAME)

debug:
	go build -v -gcflags=all="-N -l" -o $(NAME) ./cmd/$(NAME)

# Quality targets
fmt:
	go fmt ./...

lint:
	golangci-lint run


# Test targets
test:
	go test ./...

test-parsing:
	@./scripts/prs-tests.sh

test-instantiating:
	@./scripts/ins-tests.sh

test-registerification:
	@./scripts/reg-tests.sh

test-expr:
	@./scripts/expr-tests.sh

test-all: test test-parsing test-expr test-instantiating test-registerification


# Installation targets
install:
	cp $(NAME) /usr/bin

uninstall:
	rm /usr/bin/$(NAME)
