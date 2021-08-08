package main

import (
	"fmt"
	"strconv"
)

func main()  {
	// Atoi() : 字符串 转 整数
	s1 := "2312"
	n1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", n1, n1)	// type:int value:2312
	}
	
	// Itoa() : 整数 转 字符串
	n2 := 200
	s2 := strconv.Itoa(n2)
	fmt.Printf("type:%T value:%#v\n", s2, s2)	// type:string value:"200"

	// Parse系列函数：用于转换字符串为给定类型的值
	b, err := strconv.ParseBool("true")
	fmt.Println(b)

	i, err := strconv.ParseInt("2485", 10, 32)
	fmt.Println(i)

	u, err := strconv.ParseUint("3271", 10, 16)
	fmt.Println(u)

	f, err := strconv.ParseFloat("3.14159", 64)
	fmt.Println(f)

	// Format系列函数：实现了将给定类型数据格式化为string类型数据的功能。
}
