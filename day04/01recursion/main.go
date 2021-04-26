package main

import "fmt"

// 递归

// 阶乘
func f(n int) int {
	if n == 1 {
		return 1
	}
	return n * f(n - 1)
}

func main()  {
	fmt.Println(f(5))
}
