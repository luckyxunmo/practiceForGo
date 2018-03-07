package main

import (
	"fmt"
	"sync"
)

var c = make(chan int, 50)

func main() {
	wg :=  sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go consumer(i,&wg)
	}
	for i := 0; i < 1000; i++ {
		c <- i
	}
	close(c)
   wg.Wait()
}

func consumer(index int,wg *sync.WaitGroup ) {
	for target := range c {
		fmt.Printf("no.%d:%d\n", index, target)

	}
	defer wg.Done()
}
