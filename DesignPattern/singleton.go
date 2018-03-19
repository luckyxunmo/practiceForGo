package design

import "sync"

type singleton struct{
	data int
}
var sin *singleton
var once sync.Once

func GetSingleton() *singleton  {
	once.Do(func(){
		sin = &singleton{}
	})
	return sin
}

