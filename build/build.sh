#!/bin/bash
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}")" && pwd )
ROOT=$DIR/..
GO_PATH=$DIR/../../../../..
export GOPATH=$GO_PATH
go build $ROOT/main.go
