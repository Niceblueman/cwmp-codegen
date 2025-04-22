.PHONY: build test clean bench install release

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
BINARY_NAME=cwmp-codegen
BINARY_UNIX=$(BINARY_NAME)_unix
MAIN_PATH=./cmd/cwmp-codegen

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_PATH)

test:
	$(GOTEST) -v ./...

coverage:
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out

bench:
	$(GOTEST) -bench=. ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -f coverage.out
dev:
	$(GOCMD) run $(MAIN_PATH) $(ARGS)
tr069-test:
	$(GOCMD) run $(MAIN_PATH) -input https://raw.githubusercontent.com/BroadbandForum/cwmp-data-models/refs/heads/master/tr-069-1-0-0-full.xml -output output/tr069
run:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_PATH)
	./$(BINARY_NAME)

install:
	$(GOINSTALL) $(MAIN_PATH)

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v $(MAIN_PATH)

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME).exe -v $(MAIN_PATH)

build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)_mac -v $(MAIN_PATH)

release: build-linux build-windows build-mac