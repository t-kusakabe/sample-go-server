##########################
## Builder Container
##########################
FROM golang:1.18.1-alpine3.15 as builder

WORKDIR /go/src/github.com/sample-go-server

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN apk update && \
      apk add --no-cache git && \
      go build -o app main.go

##########################
## Application Container
##########################
FROM alpine

WORKDIR /sample-go-server

RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/github.com/sample-go-server /sample-go-server

CMD ["./app"]
