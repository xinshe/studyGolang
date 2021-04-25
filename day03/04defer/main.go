package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

// defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}

/*
A 1 2 3
B 10 2 12
BB 10 12 22
AA 1 3 4
*/
