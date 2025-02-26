
FROM golang:1.20


WORKDIR /app


COPY tcp_server.go .


RUN go build -o tcp_server tcp_server.go


CMD ["./tcp_server"]