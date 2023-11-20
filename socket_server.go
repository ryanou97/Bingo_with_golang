package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		// 讀取客戶端發送的資料
		n, err := conn.Read(buffer)

		if err != nil {
			// 加入這段，可以防止client斷線後print出 "Error reading: EOF"
			if err.Error() == "EOF" {
				fmt.Println("Client closed the connection")
				return
			}
			fmt.Println("Error reading:", err)
			return
		}

		// 將資料轉換成字串並顯示
		message := string(buffer[:n])
		fmt.Println("Received message:", message)

		// 回覆客戶端
		conn.Write([]byte("Message received"))
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}

		fmt.Println("Accepted connection from:", conn.RemoteAddr())

		// 啟動一個 goroutine 處理連線
		go handleConnection(conn)
	}
}
