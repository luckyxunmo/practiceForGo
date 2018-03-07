package main

import (
	"sync"
	"time"
	"fmt"
)

func timeChan(wg *sync.WaitGroup)  {
	tick := time.Tick(100*time.Microsecond)
	boom := time.After(5000*time.Microsecond)
	defer wg.Done()
	for{
		select{
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("end")
			return
		default:
			fmt.Println(".")
			time.Sleep(50*time.Microsecond)

		}
	}
}
func main(){
	var wg sync.WaitGroup
	wg.Add(1)
	go timeChan(&wg)
	wg.Wait()
}
