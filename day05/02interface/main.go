package main

import (
	"fmt"
)

// 接口示例2

type car interface {
	run()
}

type falali struct {
	color string
}

func (f falali) run()  {
	fmt.Printf("%s的法拉利\n", f.color)
}

type baoshijie struct {
	color string
}

func (b baoshijie) run()  {
	fmt.Printf("%s的保时捷", b.color)
}

func drive(c car)  {
	c.run()
}

func main() {
	f := &falali{
		color: "红色",
	}

	b := &baoshijie{
		color: "黑色",
	}

	drive(f)
	drive(b)
}
