package main

import "fmt"

// 值接收者和指针接收者实现接口的区别

type speaker interface {
	speak()
}

type dog struct {
}

// 值接收者实现接口
func (d dog) speak()  {
	fmt.Println("汪汪汪")
}

type cat struct {
}

// 指针接收者实现接口
func (c *cat) speak()  {
	fmt.Println("喵喵喵")
}

/*
使用值接收者实现接口之后，不管是dog结构体还是结构体指针*dog类型的变量都可以赋值给该接口变量。
因为Go语言中有对指针类型变量求值的语法糖，dog指针d2内部会自动求值*d2。

总之，
使用值接收者实现接口，结构体类型和结构体指针类型 的变量 都可以保存在 接口变量中
指针接收者实现接口，接口变量中 只能保存 结构体指针类型的变量
*/
func main()  {
	var x speaker

	d1 := dog{}
	x = d1
	x.speak()

	d2 := &dog{}
	x = d2
	d2.speak()

	//c1 := cat{}
	//x = c1	// 报错
	//c1.speak()

	c2 := &cat{}
	x = c2
	c2.speak()
}
