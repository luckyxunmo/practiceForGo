package main

import (
	"net"
	"fmt"
	"bufio"
)

func main()  {
	l,err := net.Listen("tcp","127.0.0.1:8053")
	if err != nil{
		fmt.Printf("Fail to listen:%s\n",err.Error())
	}
	for{
		if c,err := l.Accept();err == nil{
		      go Echo(c)
		}
	}
}
func Echo(c net.Conn)  {
	defer c.Close()
	line,err :=bufio.NewReader(c).ReadString('\n')
	if err != nil{
		fmt.Printf("fail to read:%s\n",err.Error())
		return 
	}
	if _,err := c.Write([]byte(line));err != nil{
		fmt.Printf("fail to write:%s\n",err.Error())
		return
	}
}

