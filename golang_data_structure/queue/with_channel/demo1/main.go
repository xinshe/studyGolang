package main

import (
	"fmt"
)

// 用channel实现一个定长队列（测试版）
func writeQueue(queue chan int64, data int64) {
	select {
	case queue <- data:
		fmt.Println("case push ", data, "\t len=", len(queue))
	default:
		pollData := <-queue
		fmt.Println("poll ", pollData, "\t len=", len(queue))
		queue <- data
		fmt.Println("default push ", data, "\t len=", len(queue))
	}
}

func main() {
	queue := make(chan int64, 10)
	for i := 1; i <= 1000; i++ {
		//data := rand.Int63n(100)
		writeQueue(queue, int64(i))
	}
	fmt.Println(len(queue))
}
