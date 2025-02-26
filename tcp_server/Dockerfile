#Docker konteynerinde Go derleyicisinin mevcut olmasını sağlar.
FROM golang:1.20

#Docker konteynerinde /app dizinine geçiş yapar ve tüm sonraki komutların bu dizin içinde çalıştırılmasını sağlar.
WORKDIR /app

#Yerel bilgisayarınızda bulunan tcp_server.go dosyasını Docker konteynerindeki /app dizinine kopyalar.
COPY tcp_server.go .

#Bu satır, tcp_server.go dosyasını derleyerek tcp_server adında bir çalıştırılabilir dosya oluşturur.
RUN go build -o tcp_server tcp_server.go

#Docker konteyneri başlatıldığında çalıştırılacak varsayılan komutu belirtir.
CMD ["./tcp_server"]
