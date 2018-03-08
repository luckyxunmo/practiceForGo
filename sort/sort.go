package main

import (
	"sort"
	"fmt"
)

type student struct{
	age int
}
type StuVec []student
func (a StuVec) Len() int  {
	return len(a)
}
// 从大到小
func (a StuVec)Less(i,j int) bool {
	return a[i].age > a[j].age
}
func (a StuVec)Swap(i,j int)  {
	a[i],a[j] = a[j],a[i]
}
func main()  {
	b := student{
		age:2,
	}
	c := student{
		age:1,
	}
	d := student{
		age:3,
	}
	e := student{
		age:6,
	}
	vec := StuVec{}
	vec = append(vec, b,c,d,e)
	sort.Sort(vec)
	fmt.Println(vec)
}
