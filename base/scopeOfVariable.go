/*
对于使用:=定义的变量，如果新变量p与那个同名已定义变量 (这里就是那个第9行的局部变量p)不在一个作用域中时，
那么golang会新定义这个变量p，遮盖住全局变量p，
 第1个的p 与第2个的p不是同一个变量，两者不是同一个作用域
*/

package main

import "fmt"

func main() {
	var p int // 1

	if 3 < 5 {
		p, err := testScope() //2
		if err == nil {
			fmt.Println(p) // p is 3
		}
	} else {
		fmt.Println("D")
	}
	fmt.Println("p is:", p) // p is 0
}
func testScope() (int, error) {
	return 3, nil
}
