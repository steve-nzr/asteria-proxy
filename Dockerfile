FROM golang:1.12-alpine as pre-builder
ENV GO111MODULE on
ENV GOPATH /go
WORKDIR /go/src/github.com/steve-nzr/asteria-proxy

RUN apk add git gcc g++

COPY . .

RUN go mod download
RUN go get github.com/cortesi/modd/cmd/modd

ENTRYPOINT [ "modd", "-f", "configs/modd/app.conf" ]

FROM pre-builder as builder
RUN go build -o /bin/app cmd/app/main.go

FROM alpine:3.10 as runtime
RUN apk add ca-certificates

COPY --from=builder /bin/app /bin/
ENTRYPOINT [ "/bin/app" ]
