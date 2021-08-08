package main

import "fmt"

func main()  {
	ch := make(chan int, 2)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
			fmt.Println("发送成功")
		case x := <-ch:
			fmt.Println(x)
		}
	}
}
