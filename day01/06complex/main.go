package main

import "fmt"

// 复数
// 复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位。

func main() {
	c1 := 3 + 2i
	fmt.Println(c1)         // (3+2i)
	fmt.Printf("%T \n", c1) // complex128	默认为complex128类型

	// 显式声明为complex64
	c2 := complex64(1 + 2i)
	fmt.Println(c2)
	fmt.Printf("%T \n", c2)

	var c3 complex64
	fmt.Println(c3) // (0+0i)
}
