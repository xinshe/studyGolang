package main

import "fmt"

// 空接口

/*
空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
空接口类型的变量可以存储任意类型的变量。

interface : 接口关键字
interface{} : 空接口类型

空接口的应用：
1. 空接口作为函数的参数
2. 空接口作为map的值
 */

// 空接口作为函数的参数
// 使用空接口实现可以接收任意类型的函数参数。
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main()  {
	// 定义一个空接口类型的变量
	var x interface{}
	s := "hello"
	x = s
	fmt.Printf("type:%T, value:%v\n", x, x)

	i := int64(235)
	x = i
	fmt.Printf("type:%T, value:%v\n", x, x)

	b := false
	x = b
	fmt.Printf("type:%T, value:%v\n", x, x)

	// 空接口作为函数的参数
	show(123)
	show(false)
	show(nil)

	// 空接口作为map的值
	// 使用空接口实现可以保存任意值的字典。
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "沙河娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)
}

