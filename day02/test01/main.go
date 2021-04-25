package main

import "fmt"

/*
有一堆数字，如果除了一个数字以外，其他数字都出现了两次，那么如何找到出现一次的数字？
*/

func main() {
	var arr = [9]int{2, 4, 5, 7, 8, 5, 7, 4, 8}
	var res int
	switch len(arr) {
	case 0:
		fmt.Println("...")
		return
	case 1:
		fmt.Println(arr[0])
		return
	default:
		res = arr[0] ^ arr[1]
	}

	for i := 2; i < len(arr); i++ {
		res ^= arr[i]
	}
	fmt.Println(res)
}
