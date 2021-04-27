package main

import "fmt"

// 结构体初始化

type person struct {
	name string
	age int
	gender string
	hobby []string
}

func main()  {

	// 1. 没有初始化的结构体，其成员变量都是对应其类型的零值。
	p := person{}
	fmt.Printf("%T\n", p)	// main.person
	fmt.Printf("%v\n", p)	// { 0  []}
	fmt.Printf("%#v\n", p)	// main.person{name:"", age:0, gender:"", hobby:[]string(nil)

	// 2. 使用键值对初始化
	/*
	使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值。
	 */
	p1 := person{
		name: "Andy",
		age: 23,
		gender: "male",
		hobby: []string{"羽毛球"},	// 这个逗号也要加
	}
	fmt.Printf("%T\n", p1)	// main.person
	fmt.Printf("%v\n", p1)	// {Andy 23 male [羽毛球]}
	fmt.Printf("%#v\n", p1)	// main.person{name:"Andy", age:23, gender:"male", hobby:[]string{"羽毛球"}}

	/*
	对结构体指针进行键值对初始化
	注意：当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。
	 */
	p2 := &person{
		name:   "Bobo",
		gender: "female",
	}
	fmt.Printf("%T\n", p2)	// *main.person
	fmt.Printf("%v\n", p2)	// &{Bobo 0 female []}
	fmt.Printf("%#v\n", p2)	// &main.person{name:"Bobo", age:0, gender:"female", hobby:[]string(nil)}

	// 3. 使用值的列表初始化
	/*
	初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值。

	使用这种格式初始化时，需要注意：
		必须初始化结构体的所有字段。
		初始值的填充顺序必须与字段在结构体中的声明顺序一致。
		该方式不能和键值初始化方式混用。
	 */
	p3 := person{
		"Sam",
		23,
		"female",
		[]string{"钢琴"},
	}
	fmt.Printf("%T\n", p3)	// main.person
	fmt.Printf("%v\n", p3)	// {Sam 23 female [钢琴]}}
	fmt.Printf("%#v\n", p3)	// main.person{name:"Sam", age:23, gender:"female", hobby:[]string{"钢琴"}}


}
