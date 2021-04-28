package main

import "fmt"

// 任意类型添加方法
/*
在Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。
举个例子，我们基于内置的int类型使用type关键字可以定义新的自定义类型，然后为我们的自定义类型添加方法。

注意事项： 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。
 */

type MyInt int

// 给自定义类型加方法
func (i MyInt) hello()  {
	fmt.Println("hello ", i)
}

func main()  {
	n := MyInt(127)
	n.hello()
}
