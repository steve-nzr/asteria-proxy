FROM golang:1.12-alpine as builder
ENV GO111MODULE on
ENV GOPATH /go
WORKDIR /go/src/github.com/steve-nzr/asteria-proxy

RUN apk add git gcc g++

COPY . .

RUN go mod download
RUN go get github.com/cortesi/modd/cmd/modd

ENTRYPOINT [ "modd", "-f", "configs/modd/app.conf" ]
