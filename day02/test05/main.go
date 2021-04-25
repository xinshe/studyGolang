package main

import (
	"fmt"
	"sort"
)

// 请使用内置的sort包对数组var a = [...]int{3, 7, 8, 9, 1}进行排序

func main() {
	var a = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a[:])
	fmt.Println(a)
}
