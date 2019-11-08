# Makefile
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_BUILD_RACE=$(GO_CMD) build -race
GO_TEST=$(GO_CMD) test
GO_TEST_VERBOSE=$(GO_CMD) test -v
GO_TEST_COVER=$(GO_CMD) test -cover
GO_INSTALL=$(GO_CMD) install -v
GO_DEP_DOWNLOAD=$(GO_CMD) mod download
GO_BUILD_FLAGS=-ldflags "-w -s"
# GOOS=linux

SERVER_BIN=angago
SERVER_DIR=.
SERVER_MAIN=main.go

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

all: deps test build-server run

deps:
	@echo "==>Installing dependencies ...";
	$(GO_DEP_DOWNLOAD)

test: # run tests
	@echo "==> Running tests ...";
	@$(GO_TEST_COVER) $(SOURCEDIR)/...

build-server: # build serevr
	@echo "==> Building server for ..." ${GOOS};
	@export CGO_ENABLED=0
	@GOOS=$(GOOS) $(GO_BUILD) -o $(SERVER_BIN) $(GO_BUILD_FLAGS) || exit 1;
	@chmod 755 $(SERVER_BIN)

run: # run server
	@echo "==> Running server ...";
	./$(SERVER_BIN)