package main

import "fmt"

// for循环结构

func main() {
	// 基本结构
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 变种1
	var j = 5
	for ; j < 10; j++ {
		fmt.Println(j)
	}

	// 变种2 (类似while)
	k := 3
	for k < 6 {
		fmt.Println(k)
		k++
	}

	// 无限循环
	//for {
	//	fmt.Println("123")
	//}

	// for range
	/*
		Go语言中可以使用for range遍历数组、切片、字符串、map 及通道（channel）。 通过for range遍历的返回值有以下规律：
			数组、切片、字符串返回索引和值。
			map返回键和值。
			通道（channel）只返回通道内的值。
	*/
	s := "hello小明"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}

	// 流程控制之跳出for循环
	for i := 0; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	// 退出当前循环
	for i := 0; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

}
