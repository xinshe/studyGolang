package main

import "fmt"

// Go语言中使用结构体也可以实现其他编程语言中面向对象的继承。

/*
结构体字段的可见性:
结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。
 */

type animal struct {
	name string
	age int
}

func (a animal) move()  {
	fmt.Println(a.name, "在动！")
}

type dog struct {
	color string
	animal
}

func (d dog) wang()  {
	fmt.Println(d.name, "汪！")
}

func main()  {
	d1 := &dog{
		color: "black",
		animal: animal{
			name: "小明",
			age:  5,
		},
	}
	d1.move()
	d1.wang()
}
