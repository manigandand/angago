#! /bin/bash

SERVER_BIN=angago
GO_OS=darwin

# make build-server
make GOOS=$GO_OS deps test build-server

# run server
echo "==> Running server ..."
./$SERVER_BIN