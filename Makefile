 # Go parameters
    GOCMD=go
    GOBUILD=$(GOCMD) build
    GOCLEAN=$(GOCMD) clean
    GOTEST=$(GOCMD) test
    GOGET=$(GOCMD) get
    BINARY_NAME=gosearch
    BINARY_UNIX=$(BINARY_NAME)_unix


    all: test build
    build:
			$(GOBUILD) -o bin/$(BINARY_NAME) ./cmd/...
    test: 
			$(GOTEST) -v ./...
    clean: 
			$(GOCLEAN)
			rm -f $(BINARY_NAME)
			rm -f $(BINARY_UNIX)
    run:
			$(GOBUILD) -o bin/$(BINARY_NAME) ./cmd/...
			./bin/$(BINARY_NAME)
    # Cross compilation
    build-linux:
			CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_UNIX) ./cmd/...