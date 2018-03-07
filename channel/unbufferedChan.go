
package main

import (
	"fmt"
)

//打印斐波那契数列
func FibonacciS(c,quit chan int){
	x,y := 0,1
	for{
		select {
		case c<-x:
			x,y = y,x+y
		case <-quit:
			return
		}
	}
}

func main()  {
	ch := make(chan int)
	quit := make(chan int)
	go func(){
		for i:=0;i<10;i++{
			fmt.Println(<-ch)
		}
		quit<-0
	}()
	FibonacciS(ch,quit)
}
