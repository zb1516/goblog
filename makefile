BINARY_NAME="goblog"
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

default:
	@$(GOBUILD) -o $(BINARY_NAME)
test:
	@$(GOTEST) -0 $(BINARY_NAME) -v ./...
run:
	@$(GOBUILD) -o $(BINARY_NAME) -v
install:
	@govendor sync -v
clean:
	@$(GOCLEAN) -v
.PHONY: default  install test  clean run


