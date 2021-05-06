package main

import "fmt"

// 接口

/*
接口（interface）定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节。

在Go语言中接口（interface）是一种类型，一种抽象的类型。
interface是一组method的集合，是duck-type programming的一种体现。接口做的事情就像是定义一个协议（规则），
只要一台机器有洗衣服和甩干的功能，我就称它为洗衣机。不关心属性（数据），只关心行为（方法）。

 */

// 定义一个能叫的类型
type speaker interface {
	speak() // 只要实现了 speak方法的变量都是 speaker类型，方法签名
	//move()
}

type dog struct {
}

type cat struct {
}

func (d dog) speak()  {
	fmt.Println("汪汪汪")
}

func (c cat) speak()  {
	fmt.Println("喵喵喵")
}

func main()  {
	d1 := &dog{}
	c1 := &cat{}

	d1.speak()
	c1.speak()

	var s1 speaker
	s1 = d1
	s1.speak()
}
