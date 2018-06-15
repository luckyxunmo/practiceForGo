package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func signalListen(stop chan bool) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGINT)
	<-c
	fmt.Println("aaaaaaa")
	stop <- true
}

func main() {
	stop := make(chan bool)

	http.HandleFunc("/", hello)
	server := &http.Server{
		Addr:    ":1234",
		Handler: http.DefaultServeMux,
	}
	go signalListen(stop)
	go server.ListenAndServe()
	go func() {
		fmt.Println(<-stop)
		server.Close()
	}()

	time.Sleep(10 * time.Second)

}
func hello(w http.ResponseWriter, resp *http.Request) {
	writeResponse(w, nil, nil)
}
func writeResponse(resp http.ResponseWriter, err error, data interface{}) {

	fmt.Fprint(resp, "hello world lalala")
}
