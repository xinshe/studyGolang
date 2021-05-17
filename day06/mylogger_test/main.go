package main

import (
	"studyGolang/day06/mylogger"
	"time"
)

func main()  {
	log := mylogger.NewLogger("debug")
	for  {
		log.Debug("this is debug logger.")
		log.Trace("this is trace logger.")
		log.Info("this is info logger.")
		log.Warning("this is warning logger.")
		log.Error("this is error logger.")
		log.Fatal("this is fatal logger.")
		time.Sleep(time.Second * 2)
	}
}
