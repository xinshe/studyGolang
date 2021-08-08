package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func c() {
	for i := 1; i < 10; i++ {
		fmt.Println("C:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(4)
	go a()
	go b()
	go c()
	time.Sleep(time.Second)
}
