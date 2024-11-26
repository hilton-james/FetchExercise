NAME = receipt
COMMAND = server

format:
	@go fmt ./...

run:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs=false -trimpath -ldflags="-w -s" -o bin/${NAME} cmd/${NAME}/main.go
	@bin/receipt ${COMMAND}

clean: 
	@go clean -cache
