SOURCEDIR=.

GOCMD=go
GOBUILD=$(GOCMD) build
BINARY=git-releaser
BINARY_PATH=$(SOURCEDIR)/cmd/git_releaser/main.go

.DEFAULT_GOAL: $(BINARY)

all: clean build

build: 
	$(GOBUILD) -o ${BINARY} $(BINARY_PATH)

clean: 
	rm -f $(BINARY)
