package main

import "fmt"

// 从切片中删除元素
// 总结一下就是：要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)

func main() {
	// Go语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素。 代码如下：
	s := []int{30, 31, 32, 33, 34, 35, 36, 37}
	s = append(s[:2], s[3:]...)
	fmt.Println(s)      // [30 31 33 34 35 36 37]
	fmt.Println(len(s)) // 7
	fmt.Println(cap(s)) // 8

	arr := [...]int{1, 3, 5, 7, 9, 11}
	s1 := arr[:]
	fmt.Printf("%p\n", &s1[0])
	s1 = append(s1[:3], s1[4:]...) // 修改了底层数组
	fmt.Printf("%p\n", &s1[0])
	/*
		相当于：s1[:3] ==> 切片1为[1, 3, 5]  底层数组为[1, 3, 5, 7, 9, 11]
		       s1[4:] ==> 切片2为[9, 11]	底层数组为[1, 3, 5, 7, 9, 11]
		append执行的操作就是 在原底层数组的空间上，将切片2追加到切片1上，即[1, 3, 5][9, 11][11]，
		最后这个[11]为原底层数组的值，不是切片的值。
		注意：切片仅是引用，并不保存具体值，一切操作都是在底层数组上！
	*/
	fmt.Println(s1)      // [1 3 5 9 11]
	fmt.Println(len(s1)) // 5
	fmt.Println(cap(s1)) // 6	注意！

	fmt.Println(arr)      // [1 3 5 9 11 11] 注意！
	fmt.Println(len(arr)) // 6
	fmt.Println(cap(arr)) // 6

}
