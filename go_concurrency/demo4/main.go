package main

import "fmt"

func main()  {
	var c1 chan int
	fmt.Println(c1) // <nil>
	c1 = make(chan int, 2)
	c1 <- 10
	c1 <- 11
	x := <-c1
	fmt.Println(x)

	close(c1)
	//c1 <- 12	// panic: send on closed channel
	y := <-c1
	fmt.Println(y)	// 11
	z := <-c1
	fmt.Println(z)	// 0

}
