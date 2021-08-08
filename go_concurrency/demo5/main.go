package main

import (
	"fmt"
	"time"
)

// 无缓冲的通道

func recv(ch chan int)  {
	ret := <-ch
	fmt.Println("接收成功", ret)
}

func main()  {
	ch := make(chan int)
	go func() {
		ch <- 12
		fmt.Println("发送 12")
	}()
	go recv(ch)
	time.Sleep(time.Second)
	fmt.Println("done")
}
