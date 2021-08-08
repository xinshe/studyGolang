package main

import (
	"fmt"
	"time"
)

// 启动单个goroutine

func hello()  {
	fmt.Println("hello goroutine")
}

func main()  {
	go hello()
	fmt.Println("main goroutine")
	time.Sleep(time.Second)
}
