package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int{
	return len(h)
}
func (h IntHeap)Less(i,j int) bool  {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i,j int)  {
	h[i],h[j] = h[j],h[i]
}

func (h *IntHeap)Push(x interface{})  {
	*h = append(*h,x.(int))
}
func (h *IntHeap) Pop() interface{}  {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

func main()  {
	h := &IntHeap{2,17,1,57,34,67,8,9}
	heap.Init(h)
	fmt.Println("h is ",*h)
	heap.Push(h,3)
	heap.Fix(h,3)
	for h.Len() > 0 {
		fmt.Printf("%d ",heap.Pop(h))
	}
}
