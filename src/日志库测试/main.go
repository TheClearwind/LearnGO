package main

import (
	"mylog"
	"time"
)

func main() {
	log := mylog.NewFileLoger("Debug", "./", "王兰花.log", 5)
	for {
		//id:=150
		//name:="王兰花"
		log.Debug("这是一条Debug id:%d name:%s", 150, "name")
		log.Info("这是一条Info")
		log.Error("这是一条Error")
		log.Fatal("这是一条Fatal")
		time.Sleep(time.Second)
	}
}
