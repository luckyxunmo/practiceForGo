package main

import "fmt"

// 闭包访问的是参数的地址
func Fibonacci() func() int {
	back1, back2 := 0, 1
	return func() int {
		// 重新赋值
		back1, back2 = back2, (back1 + back2)
		return back1
	}
}
func adder() func(int) int {
	sum := 10
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos := adder()
	for i := 0; i < 2; i++ {
		fmt.Println(pos(i))
	}

	f := Fibonacci()
	for i := 0; i < 10; i++{
		fmt.Println(f())
	}

}

