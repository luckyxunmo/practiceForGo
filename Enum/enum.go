package main

import "fmt"

type state int
const(
	Success state = iota
	Failed
	z = "zz"
	k
	p
	u = iota
)

func main()  {
	result := Failed
	var result2 state = Success
	fmt.Println("state",result)
	fmt.Println("state2",result2)
	fmt.Printf("z:%s,k:%s,p:%s,u:%d",z,k,p,u)
}
/*
state 1
state2 0
z:zz,k:zz,p:zz,u:5
*/
