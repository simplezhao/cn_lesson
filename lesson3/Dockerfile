FROM golang:1.17-stretch as builder
WORKDIR /httpserver
COPY http_server.go .
COPY go.mod .
RUN go build -o httpserver

FROM debian:stretch as prod
WORKDIR /root
COPY --from=0 /httpserver/httpserver .
CMD ["./httpserver"]

