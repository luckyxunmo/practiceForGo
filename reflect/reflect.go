package main

import (
	"fmt"
	"reflect"
	"strconv"
)
type T struct{
	A int
	B string
}

func (t T)Print()  {
	fmt.Println("call method")
}

type MyType struct {
	i int
	name string
}

func (mt *MyType) SetI(i int) {
	mt.i = i
}

func (mt *MyType) SetName(name string) {
	mt.name = name
}

func (mt *MyType) String() string {
	return fmt.Sprintf("%p",mt) + "--name:" + mt.name + " i:" + strconv.Itoa(mt.i)
}
func main(){
	simpleReflect()
	//callFunction()
}
func simpleReflect(){
	var x float64 = 3.4
	fmt.Println("type",reflect.TypeOf(x))
	v := reflect.ValueOf(x)

	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:",v.Kind() == reflect.Float64)
	fmt.Println("value:",v.Float())

	t := T{23,"tet"}
	s := reflect.ValueOf(&t).Elem()
    s.MethodByName("Print").Call(nil)
	typeofT := s.Type()
	fmt.Println(typeofT)
	for i:=0; i < s.NumField();i++{
		f := s.Field(i)
		fmt.Printf("%d:%s %s = %v\n",i,typeofT.Field(i).Name,f.Type(),f.Interface())
	}
}

func callFunction(){
	myType := &MyType{22,"wowzai"}
	mtV := reflect.ValueOf(&myType).Elem()
	fmt.Println("Before:",mtV.MethodByName("String").Call(nil)[0])
	params := make([]reflect.Value,1)
	params[0] = reflect.ValueOf(18)
	mtV.MethodByName("SetI").Call(params)
	params[0] = reflect.ValueOf("reflection test")
	mtV.MethodByName("SetName").Call(params)
	mtV.Method(1).
	fmt.Println("After:",mtV.MethodByName("String").Call(nil)[0])
}
