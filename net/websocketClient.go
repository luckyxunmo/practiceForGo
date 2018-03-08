package main

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"fmt"
)

var origin = "http://127.0.0.1:1234"
var url = "ws://127.0.0.1:1234/echo"

func main(){
	ws,err := websocket.Dial(url,"",origin)
	if err != nil {
		log.Fatal(err)
	}
	message :=[]byte("hello,你好")
	_,err = ws.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("send:",message)

	var msg = make([]byte,512)
	m,err := ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("receive:",msg[:m])
	ws.Close()
}
