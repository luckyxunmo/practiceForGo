/*
* close 一个chan的时候，消费chan的地方都会收到信号
当关闭testchan 这个chan时，doEast()和doEast2()都会收到信号
与此有异曲同工的是context
*/

package main

import (
	"log"
	"time"
)

var testChan = make(chan struct{})

func someHander2() {
	go doEast()
	time.Sleep(10 * time.Second)
	close(testChan)
}

func doEast() {
	go doEast2()
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-testChan:
			log.Println(" do stuff done")
			return
		default:
			log.Println("work")
		}
	}
}
func doEast2() {
	for {
		time.Sleep(2 * time.Second)
		select {
		case <-testChan:
			log.Println(" do stuff2 done")
			return
		default:
			log.Println("doStuff2 work")
		}
	}
}
func main() {
	someHander2()
	time.Sleep(25 * time.Second)
}
