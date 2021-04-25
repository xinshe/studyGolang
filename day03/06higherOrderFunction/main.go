package main

import (
	"errors"
	"fmt"
)

// 高阶函数分为函数作为参数和函数作为返回值两部分。

// 函数可以作为参数：
func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {
	res := calc(3, 6, add)
	fmt.Println(res)
	op, err := do("+")
	if err == nil {
		fmt.Println(op(2, 4))
	}
}

// 函数也可以作为返回值：
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}
