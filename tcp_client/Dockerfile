
FROM golang:1.20


WORKDIR /app


COPY tcp_client.go .


RUN go build -o tcp_client tcp_client.go


CMD ["./tcp_client"]