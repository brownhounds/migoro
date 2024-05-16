build:
	go generate ./...
	GOOS=linux go build -ldflags="-s -w" -o ./bin/migoro main.go

lint:
	golangci-lint run --fast

fix:
	golangci-lint run --fix

run_tests:
	go clean -testcache && SNAPSHOTS_DIR=./snapshots go test ./tests/... -v

run_tests_u:
	go clean -testcache && UPDATE=true SNAPSHOTS_DIR=./snapshots go test ./tests/... -v
