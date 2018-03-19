package main

import (
	"fmt"
)

func quickSort(data []int,low,high int)  {
var base int
if low < high{
	base = getPart(data,low,high)
	quickSort(data,low,base-1)
	quickSort(data,base+1,high)
}

}
func getPart(data []int,low,high int) int  {
	 temp := data[low]
	for low < high{
		for low < high && data[high] >= temp{
			high--
		}
		if low < high{
			data[low] = data[high]
		}
		for low < high && data[low] <= temp{
			low ++
		}
		if low < high{
			data[high] = data[low]
		}

	}
	data[low] = temp
	return low
}

func main(){
  data := []int{6,3,9,6,1,45,3}
  quickSort(data,0,len(data)-1)
  fmt.Println("sort data is ",data)
}
