package main

import "fmt"

/*
Go语言中目前（Go1.12）是没有异常机制，但是使用panic/recover模式来处理错误。
panic可以在任何地方引发，但recover只有在defer调用的函数中有效。

注意：
	recover()必须搭配defer使用。
	defer一定要在可能引发panic的语句之前定义。

*/
func funcA() {
	fmt.Println("func A")
}

func funcB() {
	// 此处为打开数据库连接的语句
	// 释放数据库资源
	defer func() {
		fmt.Println("释放数据库资源")
	}()
	panic("func B 出现了错误")
	fmt.Println("func B")
}

func funcC() {
	fmt.Println("func C")
}

func main() {
	funcA()
	funcB()
	funcC()
}
