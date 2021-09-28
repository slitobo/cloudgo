
local:
	@echo "build binary for mac"
	go build -o bin/httpserver main.go

build:
	@echo "build binary for linux"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/httpserver main.go

run:
	go run main.go
    