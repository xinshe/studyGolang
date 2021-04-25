package main

import "fmt"

// 求数组[1, 3, 5, 7, 8]所有元素的和

func main() {
	res := 0
	arr := [...]int{1, 3, 5, 7, 8}
	for _, v := range arr {
		res += v
	}
	fmt.Println(res)
}
