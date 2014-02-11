set GOPATH=%CD%

go install github.com/akavel/rsrc
bin\rsrc.exe -manifest tardis.exe.manifest -o src/tardis/tardis.syso

go fmt src/tardis/tardis.go
go build -ldflags="-H windowsgui" tardis
REM go build tardis

pause