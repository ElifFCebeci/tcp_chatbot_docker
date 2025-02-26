package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	serverAddr := "127.0.0.1:12345"
	// net.Dial ile TCP bağlantısı kuruldu
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Bağlantı hatası:", err)
		return
	}
	defer conn.Close()

	// Gelen mesajları dinleyen goroutine
	go func() {
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Bağlantı kesildi.")
				return
			}
			fmt.Print("Gelen mesaj: ", message)
		}
	}()

	// Kullanıcıdan mesaj alıp gönderme
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Mesaj girin: ")
		if scanner.Scan() {
			msg := scanner.Text() + "\n"
			_, err := conn.Write([]byte(msg))
			if err != nil {
				fmt.Println("Mesaj gönderme hatası:", err)
				return
			}
		} else {
			break
		}
	}
}
