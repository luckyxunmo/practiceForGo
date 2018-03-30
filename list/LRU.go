package main

import (
	"container/list"
	"fmt"
)
type Key interface{}
type entry struct{
	key Key
	value interface{}
}
/*
MaxEntries: Cache中存储的总个数,为0 则不限制
cacheMap: key为interface{}， value为list.Element其指向的类型是entry类型
*/
type Cache struct{
	MaxEntries int
	OnEvicted func(key Key, value interface{})
	ll *list.List
	cacheMap map[interface{}]*list.Element  // 保存key 和对应的链表节点
}

func New(maxEntries int) *Cache  {
	return &Cache{
		MaxEntries:maxEntries,
		ll:list.New(),
		cacheMap: make(map[interface{}] *list.Element),
	}
}

func (c *Cache) Add(key Key,value interface{})  {
	if c.cacheMap == nil{
		c.cacheMap = make(map[interface{}] *list.Element)
		c.ll = list.New()
	}
	if value,ok := c.cacheMap[key];ok{
		c.ll.MoveToFront(value)
		value.Value.(*entry).value = value
	}

	ele := c.ll.PushFront(&entry{key,value})
	c.cacheMap[key] = ele
	if c.MaxEntries != 0 && c.ll.Len() > c.MaxEntries{
		c.RemoveOldest()
	}
}

func (c *Cache)Get(key Key)(value interface{},ok bool)  {
	if c.cacheMap == nil{
		return
	}
	if value,exist := c.cacheMap[key];exist{
		c.ll.MoveToFront(value)
		return value.Value.(*entry).value,true
	}
	return
}

func (c *Cache)Delete(key Key)  {
	if c.cacheMap == nil{
		return
	}
	if value,ok := c.cacheMap[key];ok{
		c.deleteElement(value)
	}
}

// 删除最久最少未被访问的数据
func (c *Cache) RemoveOldest()  {
	if c.cacheMap == nil{
		return
	}
	ele := c.ll.Back()
	if ele != nil{
		c.deleteElement(ele)
	}
}
func (c *Cache)deleteElement(e *list.Element)  {
	c.ll.Remove(e)
	kv := e.Value.(*entry)
	delete(c.cacheMap,kv.key)
	if c.OnEvicted != nil{
		c.OnEvicted(kv.key,kv.value)
	}
}
func (c *Cache)PrintKey()  {
	for e := c.ll.Front();e != nil; e = e.Next(){
		fmt.Println(e.Value.(*entry).key)
	}
}

func main()  {
	var ca Cache
	ca.Add(123,"1")
	ca.Add(456,"2")
	ca.Add(456,"3")
	ca.Add(7,"7")
	ca.Get(456)
    ca.PrintKey()

}

