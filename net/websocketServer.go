package main

import (
	"net/http"
	"code.google.com/p/go.net/websocket"
	"log"
	"fmt"
)

func echoHandler(ws *websocket.Conn)  {
	msg := make([]byte,512)
	n,err := ws.Read(msg)
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println("Receive:%s",msg[:n])
	sendMsg := string(msg[:n])+"server recieve"
	m,err := ws.Write([]byte(sendMsg))
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println("send:%s",sendMsg[:m])
}
func main()  {
	http.Handle("/echo",websocket.Handler(echoHandler))
	http.Handle("/",http.FileServer(http.Dir(".")))

	err := http.ListenAndServe(":1234",nil)
	if err != nil{
		panic("ListenAndServe:"+err.Error())
	}
}

