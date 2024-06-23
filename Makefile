GO = go
GOBUILD = $(GO) build -mod vendor
GOTEST  = $(GO) test
GOCLEAN = $(GO) clean
APP = dups

DEBUGFLAGS = -race -gcflags="-m -l"
DEBUGBUILD = $(GO) build -mod vendor $(DEBUGFLAGS)

.PHONY: all build test clean

all: test build

build:
	$(GOBUILD) -o $(APP) main.go

test:
	$(GOTEST) -v ./...
debug:
	$(DEBUGBUILD) -o $(APP) main.go

clean:
	rm $(APP)

