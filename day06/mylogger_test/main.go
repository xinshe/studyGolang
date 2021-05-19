package main

import (
	"studyGolang/day06/mylogger"
	"time"
)

func main()  {
	//log := mylogger.NewConsoleLogger("error")
	log := mylogger.NewFileLogger("info", "day06/mylogger_test", "apple", 10*1024*1024)
	for  {
		name := "Andy"
		id := 1002
		log.Debug("this is debug logger.")
		log.Trace("this is trace logger.")
		log.Info("this is info logger.")
		log.Warning("this is warning logger.")
		log.Error("this is error logger. name:%v, id:%v", name, id)
		log.Fatal("this is fatal logger.")
		time.Sleep(time.Second * 3)
	}
}
