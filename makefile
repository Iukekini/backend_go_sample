
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=podium-backend-assessment
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_WINDOWS=$(BINARY_NAME)_win.exe
DEPCMD=dep
DEPENSURE=$(DEPCMD) ensure

all: test build
build: 
	$(DEPENSURE) 
	$(GOBUILD) -o $(BINARY_NAME) -v
test: 
	$(DEPENSURE) -v
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -f $(BINARY_WINDOWS)
run:
	$(DEPENSURE)
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
deps:
	$(GOGET) github.com/markbates/goth
	$(GOGET) github.com/markbates/pop


# Cross compilation
build-all: build-linux build-windows
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=386 $(GOBUILD) -o $(BINARY_WINDOWS) -v