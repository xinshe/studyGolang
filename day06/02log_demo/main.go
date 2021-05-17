package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.OpenFile("day06/02log_demo/xx.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	//log.SetOutput(os.Stdout)
	log.SetOutput(file)
	for  {
		log.Println("这是一条日志")
		time.Sleep(time.Second * 3)
	}
}
