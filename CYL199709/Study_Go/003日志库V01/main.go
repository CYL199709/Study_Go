package main

import (
	"time"

	"github.com/CYL199709/Study_Go/003日志库V01/mylogger"
)

//测试我们自己写的日志库
func main() {
	log := mylogger.NewFileLogger("Debug", "./", "cyl.log", 10*1024)
	for {
		id := 10010
		name := "理想"
		log.Debug("这是一条Debug日志，id:%d,name:%s", id, name)
		log.Info("这是一条info日志")
		log.Warning("这是一条Warning日志")

		log.Error("这是一条Error日志")
		log.Fatal("这是一条Fatal日志")
		time.Sleep(2 * time.Second)
	}

}
