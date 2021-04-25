package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// 输出
	// 1. Print
	/*
		Print系列函数会将内容输出到系统的标准输出，区别在于
			Print函数直接输出内容
			Printf函数支持格式化输出字符串
			Println函数会在输出内容的结尾添加一个换行符
	*/
	fmt.Print("apple")
	fmt.Println("hello")
	fmt.Printf("%d", 12)

	// 2. Fprint
	/*
		Fprint 系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容。
	*/
	//func Fprint(w io.Writer, a ...interface{}) (n int, err error)
	//func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
	//func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	file, err := os.OpenFile("day03/15fmt/sample.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err: ", err)
		return
	}
	fileContent := `这是文件的内容
第一条
第二条`
	// 注意，只要满足io.Writer接口的类型都支持写入。
	fmt.Fprintf(file, "往文件里写入：%s", fileContent)

	// 3. Sprint
	/*
	Sprint系列函数会把传入的数据生成并返回一个字符串。
		func Sprint(a ...interface{}) string
		func Sprintf(format string, a ...interface{}) string
		func Sprintln(a ...interface{}) string
	 */
	s1 := fmt.Sprint("沙河小王子")
	name := "沙河小王子"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("沙河小王子")
	fmt.Println(s1, s2, s3)

	// 4. Errorf
	/*
	Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。
		func Errorf(format string, a ...interface{}) error
	 */
	// 通常使用这种方式来自定义错误类型，例如：
	err2 := fmt.Errorf("这是一个错误")
	fmt.Println(err2)
	// Go1.13版本为fmt.Errorf函数新加了一个 %w 占位符用来生成一个可以包裹Error的 Wrapping Error。
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误%w", e)
	fmt.Println(w)


}
