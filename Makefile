build:
	go generate ./...
	GOOS=linux go build -ldflags="-s -w" -o ./bin/migoro main.go

lint:
	golangci-lint run --fast

fix:
	golangci-lint run --fix

move-version:
	cp ./VERSION ./version/VERSION

publish:
	echo $(v) > ./VERSION
	cp ./VERSION ./version/VERSION

	docker build -t brownhounds/migoro .
	docker image tag brownhounds/migoro brownhounds/migoro:$(v)
	docker push brownhounds/migoro:$(v)

	docker image tag brownhounds/migoro brownhounds/migoro:latest
	docker push brownhounds/migoro

git-tag:
	git tag --sign $(v) -m $(v)
	git push origin $(v)

install:
	GOOS=linux go build -ldflags="-s -w" -o ./bin/migoro main.go
	cp ./bin/migoro ~/.local/bin/migoro

targets:
	@go tool dist list

ci-build:
	go generate ./...
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ./build/migoro-windows-amd64.exe main.go
	GOOS=windows GOARCH=arm64 go build -ldflags="-s -w" -o ./build/migoro-windows-arm64.exe main.go

	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./build/migoro-linux-amd64 main.go
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o ./build/migoro-linux-arm64 main.go

	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ./build/migoro-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o ./build/migoro-darwin-arm64 main.go
