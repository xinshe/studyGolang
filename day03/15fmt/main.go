package main

import (
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

}
