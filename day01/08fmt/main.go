package main

import "fmt"

// fmt占位符

func main() {
	var n = 108
	fmt.Printf("%T \n", n)  // int
	fmt.Printf("%v \n", n)  // 108
	fmt.Printf("%#v \n", n) // 108
	fmt.Printf("%b \n", n)  // 1101100
	fmt.Printf("%d \n", n)  // 108
	fmt.Printf("%o \n", n)  // 154
	fmt.Printf("%x \n", n)  // 6c
	fmt.Printf("%X \n", n)  // 6C

	var s = "Hello 小明"
	fmt.Printf("%T \n", s) // string
	fmt.Printf("%s \n", s) // Hello 小明
	fmt.Printf("%v \n", s) // Hello 小明
	fmt.Printf("%#v", s)   // "Hello 小明"

}
