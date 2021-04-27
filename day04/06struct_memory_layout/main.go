package main

import (
	"fmt"
	"unsafe"
)

// 结构体内存布局

// 结构体占用一块连续的内存。
type test struct {
	a int8
	b int8
	c int8
	d string
	e int8
}

func main()  {
	t := test{
		a: 1,
		b: 4,
		c: 7,
		d: "hello world",
		e: 2,
	}
	fmt.Printf("%p\n", &t.a)	// 0xc0000a6020
	fmt.Printf("%p\n", &t.b)	// 0xc0000a6021
	fmt.Printf("%p\n", &t.c)	// 0xc0000a6022
	fmt.Printf("%p\n", &t.d)	// 0xc0000a6028
	fmt.Printf("%p\n", &t.e)	// 0xc0000a6038

	// go中的内存对齐 https://eddycjy.com/posts/go/talk/2018-12-26-go-memory-align/

	// 空结构体
	/*
	空结构体是不占用空间的。
	 */
	var v struct{}
	fmt.Println(unsafe.Sizeof(v))  // 0
}
