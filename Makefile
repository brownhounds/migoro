build:
	GOOS=linux go build -ldflags="-s -w" -o migoro main.go