package main

import "fmt"

// 使用copy()函数复制切片
/*
由于切片是引用类型，所以s1和s2其实都指向了同一块内存地址。修改s1的同时s2的值也会发生变化。
Go语言内建的copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中，copy()函数的使用格式如下：
	copy(destSlice, srcSlice []T)
其中：
	srcSlice: 数据来源切片
	destSlice: 目标切片

*/
func main() {

	s1 := []int{1, 2, 3}

	s2 := s1

	var s3 []int
	copy(s3, s1)

	var s4 = make([]int, 2)
	copy(s4, s1)

	var s5 = make([]int, 3)
	copy(s5, s1)

	var s6 = make([]int, 4)
	copy(s6, s1)

	fmt.Println(s1, s2, s3, s4, s5, s6) // [1 2 3] [1 2 3] [] [1 2] [1 2 3] [1 2 3 0]

	s1[0] = 432
	fmt.Println(s1, s2, s3, s4, s5, s6) // [432 2 3] [432 2 3] [] [1 2] [1 2 3] [1 2 3 0]

	//
}
