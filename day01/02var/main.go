package main

import "fmt"

// go语言中推荐使用驼峰式命名变量，即
var studentName string

// 还有其他命名变量的方式，如
var StudentName string
var student_name string

// 声明变量（全局变量）
var name string = "xin"
var age int
var isOK bool

// 批量声明变量（全局变量）
var (
	name1 string // ""
	age1  int    // 0
	isOK1 bool   // false
)

// go语言中，函数外的语句必须以关键字开头

func main() {
	name = "Andy"
	age = 18
	isOK = true
	fmt.Println("")

	name1 = "Jack"

	// 未使用的全局变量可以通过编译，如上面的name,age等，未使用的局部变量不能通过编译，如下面的a,b
	//var a string = "aaa"
	//var b int

	// 格式化输出
	fmt.Printf("name:%s\n", name1)

	// 声明变量时赋值
	var s1 string = "Sam"
	fmt.Println(s1)

	// 类型推导（根据值判断变量的类型）
	var s2 = 102
	fmt.Println(s2)

	// 短变量声明（在函数内部使用）
	s3 := "Teg"
	fmt.Println(s3)

	// 匿名变量（"_"）
	x, _ := foo()
	_, y := foo()
	fmt.Printf("x=%d\n", x)
	fmt.Println("y=", y)

	// 第13行声明全局变量name，下面声明了局部变量name，是可以的
	// 在同一个作用域（{}）中，不能重复声明同名的变量
	name := "she"
	fmt.Println(name)

	// 以下报错
	//var name = "Sun"
	//name := "hiroto"

}

func foo() (int, string) {
	return 10, "Andy"
}
