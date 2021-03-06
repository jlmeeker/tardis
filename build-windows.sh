#!/bin/bash

export GOPATH=`pwd`

# Build resource file
GOOS=linux go install github.com/akavel/rsrc
./bin/rsrc -manifest tardis.exe.manifest -o src/tardis/tardis.syso

# Compile Windows .exe
go fmt src/tardis/tardis.go
go fmt src/tardis/popup/*.go
go fmt src/tardis/logout/*.go
go fmt src/tardis/userinfo/*.go
GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui" -o tardis.exe tardis
