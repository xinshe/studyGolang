package main

import "fmt"

func f1() {
	fmt.Println("f1")
}

func f2() int {
	fmt.Println("f2")
	return 10
}

func f3(a, b int) int {
	fmt.Println("f3")
	return a + b
}

// 定义函数类型
// 我们可以使用type关键字来定义一个函数类型，具体格式如下：
type calculation func(int, int) int

/*
上面语句定义了一个calculation类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值。
简单来说，凡是满足这个条件的函数都是calculation类型的函数，例如上面的f3是calculation类型。
*/

func main() {
	a := f1
	fmt.Printf("%T \n", a) // func()

	b := f2
	fmt.Printf("%T \n", b) // func() int

	c := f3
	fmt.Printf("%T \n", c) // func(int, int) int

	// f3能赋值给calculation类型的变量。
	var d calculation
	d = f3
	fmt.Printf("%T \n", d) // main.calculation
}
