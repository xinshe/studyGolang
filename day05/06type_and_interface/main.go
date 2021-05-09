package main

import "fmt"

// 类型 (结构体) 与接口的关系

// 1. 一个类型（结构体）实现多个接口
// 2. 多个类型（结构体）实现同一接口

// 1. 一个类型实现多个接口
/*
一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。
*/

// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

// dog既可以实现Sayer接口，也可以实现Mover接口。
type dog struct {
	name string
}

// 实现Sayer接口
func (d dog) say() {
	fmt.Printf("%s会叫汪汪汪\n", d.name)
}

// 实现Mover接口
func (d dog) move() {
	fmt.Printf("%s会动\n", d.name)
}

func main() {
	var x Sayer
	var y Mover

	var a = dog{name: "旺财"}
	x = a
	y = a
	x.say()
	y.move()
}

// 2. 多个类型实现同一接口
// 见01

// 嵌套接口
type animal interface {
	Mover
	Sayer
}