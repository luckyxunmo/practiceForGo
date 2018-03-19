package design

import "testing"

func TestSingleton(t *testing.T)  {
	sin1 := GetSingleton()
	sin2 := GetSingleton()
	if sin1 != sin2{
		t.Error("实例对象不一样")
	}else{
		t.Log("实例对象一样")
	}
}
