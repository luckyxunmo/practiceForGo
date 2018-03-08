package main

import (
	"fmt"

	"strings"

	"strconv"
	"sync"
)

// 统计一句话中每个单词出现的次数
func WordCount(s string) map[string]int{
    var result = make(map[string]int)
	member :=strings.Split(s," ")
	for _,v := range member{
		if _,ok:=result[v];ok{
			result[v]++
		}else{
			result[v] = 1
		}
	}
	return result
}
func mapDeclare(){
	var m = map[int]int{}
	var n = make(map[int]int)
	var q = new(map[int]int) // 没有初始化
	if q == nil{
		fmt.Println("q is nil")
	}else{
		fmt.Println("q is not nil")
	}
	if *q == nil{
		fmt.Println("*q is nil")
	}else{
		fmt.Println("*q is not nil")
	}

	m[1] = 2
	n[3] = 4
	//(*q)[5] = 6 // 出错，此时 *q 为nil, q 不为nil
	*q = make(map[int]int)
	(*q)[5] = 6
	fmt.Printf("m:%v,n:%v,q:%v",m,n,q)

	if v,ok := m[1];ok{
		fmt.Println("m[1]:",v)
	}else{
		fmt.Println("NO THIS key")
	}
}
func mapCopy(){
	a = make(map[string]string)
	fmt.Println("a size is",len(a))  // size is 0, cap()不能用于map
	b=a   // a 和b 访问相同的地址
	a["hello"]="world"
	a = nil
	fmt.Println(b["hello"])
	for k := range b{
		delete(b,k)
	}

	fmt.Println("b size is",len(b))
	if b != nil{
		fmt.Println("b is not nil")
	}
	b["test"]="test"
	fmt.Println(b)
}

*
 map 的 "=" 是浅拷贝，member1和member2 指向相同的底层地址，所以test1["AA"],test1["BB"] 指向相同的内容，
*/
func memberCopy(){
	var test1 = make(map[string] map[string] string)
	member1 := map[string]string{
		"A":"A",
		"B":"b",
	}
	member2 := member1
	test1["AA"] = member1
	test1["BB"] = member2
	test1["BB"]["A"] = "good"
	fmt.Println(test1)
	// 修改了data,也就修改了里面的值（因为data 是map类型，引用传递）
	for _,data := range test1{
		//修改mem,但是没有修改原值（因为mem 是string类型，值传递）
		for _,mem := range data{
			mem = "test"
			fmt.Println(mem)
		}
		data["A"] = "bad"

	}
	fmt.Println("test 1",test1)
}


var a map[string]string
var b map[string]string

func main() {
	mapDeclare()
	mapCopy()
    testMapStruct()

	test :="I am learning Go!"
	result := WordCount(test)
	fmt.Println(result)
}

func testMapStruct()  {
	type class struct{
		name string
        sync.Mutex
	}
	xiaoMing := class{
		name:"xiaoming",
	}
	var classes = make(map[int]class)

	classes[1]=xiaoMing

	classes[2]=xiaoMing

	fmt.Println(classes[1].name)
	// v是拷贝，改变v中的内容，不会改变classes中的value
   for k,v := range classes{
  	v.Lock()
   	v.name = "tst" + strconv.Itoa(k)

   }
   // t1 是class结构体的拷贝，改变t1,不会改变map中value的值
   t1 := classes[1]
   t1.name = "123"
   fmt.Println(classes[1]) //{xiaoming {0 0}}
   // 将t1赋值给classes[1],才会改变classes[1]
   classes[1] = t1
   fmt.Println(classes[1]) //{123 {0 0}}

   // 改变classes[1]的值，不会影响改变classes[2]，因为go中赋值都是拷贝
	fmt.Println(classes) //map[1:{123 {0 0}} 2:{xiaoming {0 0}}]
}
