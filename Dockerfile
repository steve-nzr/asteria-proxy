FROM golang:1.12-alpine as builder
WORKDIR /app

RUN apk add git

COPY . .

RUN go get github.com/cortesi/modd/cmd/modd
ENTRYPOINT [ "modd", "-f", "configs/modd/app.conf" ]