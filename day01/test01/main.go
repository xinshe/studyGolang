package main

import "fmt"

/*
编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用fmt.Printf()搭配%T分别打印出上述变量的值和类型。
*/
func main() {
	var i int = 157
	var f float64 = 3.14159
	var b = true
	var s = "hello"
	fmt.Printf("i=%d (%T), f=%v (%T), b=%v (%T), s=%v (%T)", i, i, f, f, b, b, s, s)
}
