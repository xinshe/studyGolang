package main

import "fmt"

// 数组
/*
数组是存放元素的容器
必须指定存放的元素的类型和容量（长度），数组的长度是数组类型的一部分
*/
func main() {
	// 数组的声明
	var arr1 [3]bool
	var arr2 [4]int
	fmt.Printf("arr1:%T, arr2:%T \n", arr1, arr2)

	// 数组的初始化
	// 如果不初始化，默认元素都是零值（布尔值：false，整型和浮点型：0，字符串：""）
	fmt.Println(arr1, arr2)

	// 初始化方式1
	arr1 = [3]bool{false, true, true}
	fmt.Println(arr1)

	// 初始化方式2
	// 根据初始值自动推断数组的长度
	arr2 = [...]int{3, 5, 7, 8}
	fmt.Println(arr2)

	// 初始化方式3
	// 根据索引进行初始化
	arr3 := [5]int{0: 4, 3: 8}
	fmt.Println(arr3)

	// 数组的遍历
	cities := [...]string{"北京", "上海", "深圳", "长沙"}
	// 根据索引遍历
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}
	// for-range遍历
	for i, v := range cities {
		fmt.Println(i, v)
	}

	// 多维数组（以二维数组为例）
	var arr11 = [2][3]int{
		{1, 5, 7},
		{4, 7, 9},
	}
	fmt.Println(arr11)
	// 另一种写法
	var arr22 = [3][2]string{
		[2]string{"上海", "北京"},
		[2]string{"长沙", "武汉"},
		[2]string{"成都", "重庆"},
	}
	fmt.Println(arr22)

	// 多维数组的遍历
	for _, v1 := range arr22 {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 注意： 多维数组只有第一层可以使用 ... 来让编译器推导数组长度。例如：
	//支持的写法
	a := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a)

	//不支持多维数组的内层使用...
	//b := [3][...]string{
	//	{"北京", "上海"},
	//	{"广州", "深圳"},
	//	{"成都", "重庆"},
	//}
	//fmt.Println(b)

	// 数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)

	/*
		注意：
			数组支持 “==“、”!=” 操作符，因为内存总是被初始化过的。
			[n]*T表示指针数组，*[n]T表示数组指针 。
	*/

}
