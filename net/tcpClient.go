package main

import (

	"net"

	"fmt"
	"os"
	"time"
)

//
// 粘包客户端模拟
func sender(conn net.Conn) {
	for i := 0; i < 50; i++ {
		words := "{\"Id\":1,\"Name\":\"golang\",\"Message\":\"message\"}"

		conn.Write([]byte(words))
	}
}

func main() {
	server := "127.0.0.1:2343"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	fmt.Println("connect success")

	go sender(conn)

	for {
		time.Sleep(1 * 1e9)
	}
}
