#!/bin/bash

export GOPATH=`pwd`

# Compile Linux binary
go fmt src/tardis/tardis.go
go fmt src/tardis/popup/*.go
go fmt src/tardis/logout/*.go
go fmt src/tardis/userinfo/*.go
GOOS=linux GOARCH=amd64 go build tardis
