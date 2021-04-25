package main

import "fmt"

func main() {
	// new函数
	var a1 *int
	a2 := new(int)  // new函数申请一个内存地址
	fmt.Println(a1) // <nil>
	fmt.Println(a2) // 0xc000016078
	//fmt.Println(*a1)	// 报错
	fmt.Println(*a2) // 0
	*a2 = 123
	fmt.Println(*a2) // 123

	// make函数
	/*
		make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，
		而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。make函数的函数签名如下：
			func make(t Type, size ...IntegerType) Type
		make函数是无可替代的，我们在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对它们进行操作。
	*/
	var b map[string]int
	b = make(map[string]int, 10)
	b["嘻嘻哈哈"] = 100
	b["滴滴答答"] = 23
	fmt.Println(b)

	// make和new
	/*
		1. make和new都是用来申请内存的
		2. new 用来给基本数据类型（int、string）申请内存，返回的是对应类型的指针（*int、*string）
		3. make 用来给 slice、map、chan 申请内存，返回的是对应的类型的本身
	*/
}
