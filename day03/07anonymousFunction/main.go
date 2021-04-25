package main

import "fmt"

// 匿名函数多用于实现回调函数和闭包。

func main() {
	// 函数内部没有办法声明带名字的函数

	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) // 通过变量调用匿名函数

	// 如果是只调用一次的函数，还可以简写成 自执行函数
	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}
