# Go 1.20 imajını temel al
FROM golang:1.20

# Çalışma dizinini /app olarak ayarla
WORKDIR /app

# tcp_client.go dosyasını Docker konteynerine kopyala
COPY tcp_client.go .

# tcp_client.go dosyasını derle ve tcp_client adında bir çalıştırılabilir dosya oluştur
RUN go build -o tcp_client tcp_client.go

# Konteyner başladığında tcp_client çalıştırılabilir dosyasını çalıştır
CMD ["./tcp_client"]
