package main

import (
	"fmt"
	"reflect"
)
type T struct{
	A int
	B string
}
func main(){
	var x float64 = 3.4
	fmt.Println("type",reflect.TypeOf(x))
	v := reflect.ValueOf(x)

	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:",v.Kind() == reflect.Float64)
	fmt.Println("value:",v.Float())

	t := T{23,"tet"}
	s := reflect.ValueOf(&t).Elem()
	typeofT := s.Type()
	fmt.Println(typeofT)
	for i:=0; i < s.NumField();i++{
		f := s.Field(i)
		fmt.Printf("%d:%s %s = %v\n",i,typeofT.Field(i).Name,f.Type(),f.Interface())
	}
}
