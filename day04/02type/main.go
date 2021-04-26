package main

import "fmt"

// 自定义类型 和 类型别名

/*
Go语言中可以使用type关键字来定义自定义类型。
自定义类型是定义了一个全新的类型。我们可以基于内置的基本类型定义，也可以通过struct定义。例如：
将MyInt定义为int类型
	type MyInt int

通过type关键字的定义，MyInt就是一种新的类型，它具有int的特性。
 */

type NewInt int	// 自定义类型
type MyInt = int // 类型别名

func main()  {
	var a NewInt
	a = 12
	fmt.Println(a)	// 12
	fmt.Printf("%T \n", a)	// main.NewInt

	var b MyInt
	b = 13
	fmt.Println(b)	// 13
	fmt.Printf("%T\n", b)	// int

	// 内置的类型别名定义：
	/*
	type byte = uint8
	type rune = int32
	 */
	var c rune
	c = '中'
	fmt.Println(c)	// 20013
	fmt.Printf("%c \n", c)	// 中
	fmt.Printf("%T\n", c)	// int32

	/*
	结果显示a的类型是main.NewInt，表示main包下定义的NewInt类型。
	b的类型是int。MyInt类型只会在代码中存在，编译完成时并不会有MyInt类型。
	 */

}
