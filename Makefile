build:
	go generate ./...
	GOOS=linux go build -ldflags="-s -w" -o ./bin/migoro main.go

lint:
	golangci-lint run --fast

fix:
	golangci-lint run --fix
