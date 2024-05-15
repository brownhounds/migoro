build:
	go generate ./...
	GOOS=linux go build -ldflags="-s -w" -o ./bin/migoro main.go

lint:
	golangci-lint run --fast

fix:
	golangci-lint run --fix

run_tests:
	SNAPSHOTS_DIR=./snapshots go test ./tests -v

run_tests_u:
	UPDATE=true SNAPSHOTS_DIR=./snapshots go test ./tests -v
