FROM golang:alpine AS builder

LABEL org.opencontainers.image.source https://github.com/brownhounds/go-static

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/


COPY ./utils ./utils
COPY ./types ./types
COPY ./query ./query
COPY ./dispatcher ./dispatcher
COPY ./cmd ./cmd
COPY ./adapters ./adapters

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
COPY ./main.go ./main.go

RUN go get -d -v

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/migoro

FROM scratch

COPY --from=builder /go/bin/migoro /go/bin/migoro

CMD ["/go/bin/migoro", "migrate"]
