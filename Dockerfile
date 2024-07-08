FROM golang:alpine AS builder

LABEL org.opencontainers.image.source https://github.com/brownhounds/go-static

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/

COPY . .

RUN go get -d -v

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/migoro

FROM scratch

COPY --from=builder /go/bin/migoro /go/bin/migoro

ENTRYPOINT ["/go/bin/migoro"]
