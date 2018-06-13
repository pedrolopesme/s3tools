GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD)fmt
BINARY_NAME=$(GOPATH)/bin/s3tools
BINARY_UNIX=$(BINARY_NAME)_unix

build: 
	@echo "Building s3tools"
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	@echo "Running s3tools tests"
	$(GOTEST) -v ./...

clean: 
	@echo "Cleaning s3tools"
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

fmt:
	@echo "Running gofmt for all project files"
	$(GOFMT) -w */*.go
