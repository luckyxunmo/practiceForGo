package main

import (
	"fmt"
	"math"
	"io"
	"strings"
	"os"
)

//编写一个实现了 io.Reader 的 rot13Reader， 并从一个 io.Reader 读取， 利用 rot13 代换密码对数据流进行修改。
type rot13Reader struct {
	r io.Reader
}
func (reader rot13Reader)Read(b []byte)(int,error)  {
  n,err :=  reader.r.Read(b)
  for i:=0;i<n;i++{
  	switch{
	case b[i] >='A'&&b [i]<'N':
		b[i] += 13
	case b[i]>='N' && b[i] <='Z':
		b[i] -=13
	case b[i] >= 'a' && b[i] <'n' :
		b[i] += 13
	case b[i]>='n' && b[i] <='z':
		b[i] -=13
	}

  }
  return n,err
  }
type Vertex struct {
	X, Y float64
}
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
    reader := rot13Reader{s}
    io.Copy(os.Stdout,&reader)
}


