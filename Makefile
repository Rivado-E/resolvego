GOCMD := go
GOBUILD := $(GOCMD) build
GOTEST := $(GOCMD) test
GOCLEAN := $(GOCMD) clean
BINARY_NAME := resolvego

.PHONY: all
all: build

# Build the project
.PHONY: build
build:
	$(GOBUILD) -o $(BINARY_NAME) ./app

# Run tests
.PHONY: test
test:
	$(GOTEST) ./...

# Clean up binaries and cache
.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME) 

# Run the application
.PHONY: run
run:
	$(GOCMD) run ./app

.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  all         - Run tests and build the project"
	@echo "  build       - Build the project"
	@echo "  test        - Run tests"
	@echo "  clean       - Clean up binaries and cache"
	@echo "  run         - Run the application"
	@echo "  install     - Install the binary"
