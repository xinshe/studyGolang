package main

import "fmt"

func funcA() {
	fmt.Println("func A")
}

func funcB() {
	// 此处为打开数据库连接的语句
	// 释放数据库资源
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("尝试修复...")
		}
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
