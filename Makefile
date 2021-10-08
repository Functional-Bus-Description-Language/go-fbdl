PROJECT_NAME=fbdl

default: build

all: fmt vet build

build:
	go build -v -o $(PROJECT_NAME).bin .
	
help:
	@echo "Available targets:"
	@echo "  all      Run fmt vet build."
	@echo "  build    Build binary."
	@echo "  default  Run build."
	@echo "  fmt      Format files with go fmt."
	@echo "  help     Print help message."
	@echo "  vet      Examine go sources with go vet."
	@echo "  test-parsing  Run parsing tests."

fmt:
	go fmt ./...

vet:
	go vet ./...

test-parsing:
	./scripts/test-parsing.sh
