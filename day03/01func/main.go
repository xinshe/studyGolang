package main

import "fmt"

// 函数
/*
函数定义：
Go语言中定义函数使用func关键字，具体格式如下：
	func 函数名(参数)(返回值){
		函数体
	}
其中：
	函数名：由字母、数字、下划线组成。但函数名的第一个字母不能是数字。在同一个包内，函数名也称不能重名。
	参数：参数由参数变量和参数变量的类型组成，多个参数之间使用,分隔。
	返回值：返回值由返回值变量和其变量类型组成，也可以只写返回值的类型，多个返回值必须用()包裹，并用,分隔。
	函数体：实现指定功能的代码块。
*/

// 函数的定义
func sum(x int, y int) int {
	return x + y
}

// 参数

// 1. 类型简写
/*
函数的参数中如果相邻变量的类型相同，则可以省略类型
*/
func f1(x, y int) int {
	return x - y
}

// 2. 可变参数
/*
可变参数是指函数的参数数量不固定。Go语言中的可变参数通过在参数名后加...来标识。
注意：可变参数通常要作为函数的最后一个参数。
本质上，函数的可变参数是通过切片来实现的。
*/
func f2(b bool, x ...int) int {
	var res int
	if b {
		for _, v := range x {
			res += v
		}
	} else {
		res = 1
		for _, v := range x {
			res *= v
		}
	}
	return res
}

// 返回值
/*
Go语言中通过return关键字向外输出返回值。
*/

// 1. 多返回值
/*
Go语言中函数支持多返回值，函数如果有多个返回值时必须用()将所有返回值包裹起来。
*/
func f3(x int, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 2. 返回值命名
/*
函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return关键字返回。
*/
func f4(x int, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

// 3. 返回值补充
/*
当我们的一个函数返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片。
*/
func f5(s string) []int {
	if s == "" {
		return nil
	}
	return []int{1, 2, 4}
}

func main() {
	// 函数的调用
	res := sum(12, 45)
	fmt.Println(res)

	r1 := f2(true, 2, 4, 6, 8)
	r2 := f2(false, 1, 3, 1, 4)
	fmt.Println(r1)
	fmt.Println(r2)

	fmt.Println(f3(3, 1))
}
