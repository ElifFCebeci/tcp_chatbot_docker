package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

// eş zamanlı erişim için
var clients []net.Conn
var mutex sync.Mutex

func main() {
	// TCP sunucu oluşturma
	listener, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Println("Sunucu başlatılamadı:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Sunucu başlatıldı, bağlantılar bekleniyor...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Bağlantı hatası:", err)
			continue
		}

		mutex.Lock()
		clients = append(clients, conn)
		mutex.Unlock()

		fmt.Println("Yeni bir istemci bağlandı:", conn.RemoteAddr())

		go handleClient(conn)
	}
}

// İstemciden gelen mesajları dinleyen ve diğer istemcilere ileten fonksiyon
func handleClient(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("İstemci bağlantısı kesildi:", conn.RemoteAddr())

			// Bağlantısı kopan istemciyi listeden çıkar
			mutex.Lock()
			for i, c := range clients {
				if c == conn {
					clients = append(clients[:i], clients[i+1:]...)
					break
				}
			}
			mutex.Unlock()
			return
		}

		fmt.Printf("İstemciden gelen mesaj: %s", message)

		// Mesajı diğer istemcilere gönder
		broadcastMessage(message, conn)
	}
}

// Bir istemciden gelen mesajı tüm istemcilere ileten fonksiyon
func broadcastMessage(message string, sender net.Conn) {
	mutex.Lock()
	defer mutex.Unlock()

	for _, client := range clients {
		if client != sender {
			_, err := client.Write([]byte(message))
			if err != nil {
				fmt.Println("Mesaj gönderme hatası:", err)
			}
		}
	}
}
