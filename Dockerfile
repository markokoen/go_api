FROM golang:latest

WORKDIR /go/src/github.com/markokoen/go_api

ADD . .

WORKDIR /go/src/github.com/markokoen/go_api

RUN go build -o GoApi

ENTRYPOINT ./GoApi

EXPOSE 8080

EXPOSE 5432