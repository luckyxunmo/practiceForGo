package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func someHander() {
	ctx, cancel := context.WithCancel(context.Background())
	go doStuff(ctx)
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(10 * time.Second)
}

func doStuff(ctx context.Context) {
	ctx2 := context.WithValue(ctx, "test", "test") // 后两个参数不应该是string 或者任何内建的类型，而应该是自定义的类型
	go doStuff2(ctx2)
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			log.Println(" do stuff done")
			return
		default:
			log.Println("work")
		}
	}
}
func doStuff2(ctx context.Context) {
	fmt.Println(ctx.Value("test"))
	for {
		time.Sleep(2 * time.Second)
		select {
		case <-ctx.Done():
			log.Println(" do stuff2 done")
			return
		default:
			log.Println("doStuff2 work")
		}
	}
}
func main() {
	someHander()
	time.Sleep(25 * time.Second)
}
