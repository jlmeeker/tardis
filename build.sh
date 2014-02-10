#!/bin/bash

export GOPATH=`pwd`
go fmt src/tardis/tardis.go
GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui" -o tardis.exe tardis
