set GOPATH=%CD%

go fmt src/tardis/tardis.go
go build -ldflags="-H windowsgui" tardis

pause