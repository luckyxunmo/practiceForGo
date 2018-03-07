package main

import (
	"flag"
	"github.com/golang/glog"

	"time"
)

func main() {
	flag.Parse()    // 1
	flag.Lookup("logtostderr").Value.Set("true") //打印输入到屏幕
	flag.Lookup("log_dir").Value.Set("./") // 设置log文件的路径

	glog.Info("This is a Info log")         // 2
	glog.Warning("This is a Warning log")
	glog.Error("This is a Error log")

	glog.V(1).Infoln("level 1")     // 3
	glog.V(2).Infoln("level 2")
	
	glog.Flush()    // 4

	time.Sleep(100*time.Second)
}


