package main

import "fmt"

// 嵌套结构体

type address struct {
	city string
	province string
}

type school struct {
	name string
	level string
}

type person struct {
	name string
	age int
	addr address
	school // 嵌套匿名字段
}

func main() {
	p1 := person{
		name: "Andy",
		age:  24,
		addr: address{
			city:    "长沙",
			province: "湖南",
		},
	}
	// 当访问结构体成员时会先在结构体中查找该字段，找不到再去嵌套的匿名字段中查找。
	p1.school.name = "复旦大学"	// 匿名字段默认使用类型名作为字段名
	p1.level = "985"	// 匿名字段可以省略
	fmt.Println(p1)
	fmt.Println(p1.name, p1.addr.city)

}
