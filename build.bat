@ECHO OFF
go generate
ECHO.
go build -v
ECHO.
go test