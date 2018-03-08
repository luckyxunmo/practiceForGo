package main

import (
	"time"
	"log"
	"net/http"
	"runtime"
	"bytes"
	"strconv"
)
const (
	numPollers     = 2                // number of Poller goroutines to launch // Poller Go程的启动数
	pollInterval   = 60 * time.Second // how often to poll each URL            // 轮询每一个URL的频率
	statusInterval = 10 * time.Second // how often to log status to stdout     // 将状态记录到标准输出的频率
	errTimeout     = 10 * time.Second // back-off timeout on error             // 回退超时的错误
)

var urls = []string{
	"http://www.baidu.com/",
	"https://www.jd.com/2017?t=2",
	"http://blog.golang.org/",
}
type State struct{
	url string
	status string
}

type Resource struct{
	url string
	errCount int
}
func StateMonitor(updateInterval time.Duration) chan<-State{
 updates := make(chan State)
 urlStates := make(map[string]string)
 ticker := time.NewTicker(updateInterval)
 go func(){
 	for{
 		select{
 		case <-ticker.C:
 			logState(urlStates)
 			case s:= <-updates:
 				urlStates[s.url] = s.status
		}
	}
 }()
 return updates
}
func (r *Resource)Poll()string  {
	resp,err := http.Head(r.url)
	if err != nil{
		log.Println("Error",r.url,err)
		r.errCount++
		return err.Error()
	}
	r.errCount = 0
	return resp.Status
	}
func (r * Resource) Sleep(done chan<- *Resource)  {
	time.Sleep(pollInterval + errTimeout*time.Duration(r.errCount))
	done <-r
}
func logState(s map[string]string){
   log.Println("current state")
   for k,v := range s{
   	log.Printf("%s:%s",k,v)

   }
   }
func Poller(in <-chan *Resource, out chan<- *Resource,status chan<-State,i int)  {
   log.Println("process id is",i)
	for r:= range in{
		s:= r.Poll()
		status <- State{r.url,s}
		out <-r
	}

}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
func main(){
	pending,complete := make(chan *Resource),make(chan *Resource)
	status := StateMonitor(statusInterval)
	for i := 0; i < numPollers;i++{
		go Poller(pending,complete,status,i)
	}
	go func() {
		for _,url := range urls{
			pending <- &Resource{url:url}
		}
	}()
	for r:= range complete{
		go r.Sleep(pending)
	}


}
