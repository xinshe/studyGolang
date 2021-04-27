package main

import "fmt"

// 结构体实例化

/*
只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。
结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型。
	var 结构体实例 结构体类型

注意：结构体是值类型
 */

type person struct {
	name string
	age int
	gender string
	hobby []string
}

func main() {
	// 1. 基本实例化
	var p person
	p.name = "Emrys"
	p.age = 24
	p.gender = "male"
	fmt.Printf("%T\n", p)	// main.person
	fmt.Printf("%#v\n", p)	// main.person{name:"Emrys", age:24, gender:"male", hobby:[]string(nil)}

	// 2. 匿名结构体
	/*
	在定义一些临时数据结构等场景下还可以使用匿名结构体。
	 */
	var user struct{name string; age int}	// 结构体中字段如果要并排写，字段之间要加分号
	user.name = "张三丰"
	user.age = 128
	fmt.Printf("%T\n", user)	// struct { name string; age int }
	fmt.Printf("%#v\n", user) // struct { name string; age int }{name:"张三丰", age:128}

	// 3. 结构体实例化
	var p2 = person{}
	p2.name = "John"
	fmt.Printf("%T\n", p2)	// main.person
	fmt.Printf("%#v\n", p2)	// main.person{name:"John", age:0, gender:"", hobby:[]string(nil)}

	p3 := person{}
	p3.name = "Jack"

	// 4. 创建指针类型结构体
	/*
	通过使用new关键字对结构体进行实例化，得到的是结构体的地址。
	 */
	var p1 = new(person)
	(*p1).name = "Sam"
	p1.age = 28	// 语法糖，这样写在执行时会自动转成 (*p1) 的形式
	fmt.Printf("%T\n", p1)	// *main.person
	fmt.Printf("%#v\n", p1)	// &main.person{name:"Sam", age:28, gender:"", hobby:[]string(nil)}

	// 5. 取结构体的地址实例化
	/*
	使用 & 对结构体进行取地址操作相当于对 该结构体类型进行了一次new实例化操作。
	*/
	var p4 = &person{}
	p4.name = "Andy"
	(*p4).age = 18
	fmt.Printf("%T\n", p4)	// *main.person
	fmt.Printf("%#v\n", p4)	// &main.person{name:"Andy", age:18, gender:"", hobby:[]string(nil)}

}
