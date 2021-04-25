package main

import "fmt"

// defer语句
/*
Go语言中的defer语句会将其后面跟随的语句进行延迟处理。在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，
也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。

由于defer语句延迟调用的特性，所以defer语句能非常方便的处理资源释放问题。
比如：资源清理、文件关闭、解锁及记录时间、数据库连接、socket连接等。

defer后面的语句 将延迟到函数即将返回的时候再被执行
一个函数中可以有多个defer语句
多个defer语句按照先进后出的顺序延迟执行
*/

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("end")
	/*
		start
		end
		3
		2
		1
	*/
}

/*
defer执行时机：
在Go语言的函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。
而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。具体包中的图所示：
*/

// defer经典案例
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++ // 改变的是函数中x的副本
	}(x)
	return 5
}

func f5() (x int) {
	defer func(x int) int {
		x++
		return x
	}(x)
	return 5
}

func f6() (x int) {
	defer func(x *int) {
		*x++
	}(&x)
	return 5
}

func main() {
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5
	fmt.Println(f5()) // 5
	fmt.Println(f6()) // 6

}
