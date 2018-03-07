package main

import (
	"net"
	"fmt"
	"os"
)
// 粘包服务端模拟

//  conn.Read 读取指定长度的数据，即想读多少读多少字节，不会将多于给定数组的值赋值，底层应该是用缓冲
func main()  {
	netListen,err := net.Listen("tcp",":2343")
	CheckError(err)
	defer netListen.Close()
	for{
		conn,err := netListen.Accept()
		if err != nil{
			continue
		}
		go handleConnection(conn)
	}
}
func handleConnection(conn net.Conn)  {
	size := 43;
	buffer := make([]byte,size)
	for {
		var nReadLen int
		for nReadLen != size {
			var nReadLenNew int
			nReadLenNew, _ = conn.Read(buffer[nReadLen:])  //
			fmt.Println("receive data nReadLenNew length:",nReadLenNew)
			nReadLen += nReadLenNew
		}
		if nReadLen != size {
			fmt.Println("length read miss match! real=%d", len(buffer))
		}
		fmt.Println("receive data nReadLen length:",nReadLen)
		fmt.Println("receive data:",string(buffer[:nReadLen]))

	}
}
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
