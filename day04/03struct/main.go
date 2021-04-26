package main

import "fmt"

// 结构体

/*
Go语言提供了一种自定义数据类型，可以封装多个基本数据类型，这种数据类型叫结构体，英文名称struct。
也就是我们可以通过struct来定义自己的类型了。

Go语言中通过struct来实现面向对象。

结构体的定义：
使用type和struct关键字来定义结构体，具体代码格式如下：
	type 类型名 struct {
		字段名 字段类型
		字段名 字段类型
		…
	}
 */

type person struct {
	name string
	age int
	gender string
	hobby []string
}

func main()  {
	var p person
	p.name = "Andy"
	p.age = 25
	p.gender = "男"
	p.hobby = []string{"羽毛球", "狼人杀", "手游"}
	fmt.Printf("value:%v type:%T \n", p, p)	// value:{Andy 25 男 [羽毛球 狼人杀 手游]} type:main.person
	fmt.Println(p.name)
	fmt.Println(p.hobby)

}
