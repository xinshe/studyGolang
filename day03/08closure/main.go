package main

import "fmt"

// 闭包
/*
闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境。

闭包是一个函数，这个函数包含了它外部作用域的一个变量
*/

func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func adder2() func(int) int {
	var x = 10
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	ret := adder(100)

	res1 := ret(200)
	fmt.Println(res1) // 300

	res2 := ret(300)
	fmt.Println(res2) // 600

	res3 := ret(400)
	fmt.Println(res3) // 1000

	/*
		变量f是一个函数并且它引用了其外部作用域中的x变量，此时f就是一个闭包。在f的生命周期内，变量x也一直有效。
	*/
	f := adder2()
	fmt.Println(f(20)) // 30
	fmt.Println(f(40)) // 70
	fmt.Println(f(60)) // 130
}
