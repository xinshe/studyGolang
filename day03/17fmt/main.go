package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 获取输入

func main()  {
	// Go语言fmt包下有fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，可以在程序运行过程中从标准输入获取用户的输入。

	// 1. fmt.Scan
	/*
		func Scan(a ...interface{}) (n int, err error)

	Scan从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符。
	本函数返回成功扫描的数据个数和遇到的任何错误。如果读取的数据个数比提供的参数少，会返回一个错误报告原因。
	 */
	var (
		name    string
		age     int
		married bool
	)
	//fmt.Scan(&name, &age, &married)	// 因为要修改值，所以要加&
	//fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
	// fmt.Scan从标准输入中扫描用户输入的数据，将以空白符分隔的数据分别存入指定的参数。

	// 2. fmt.Scanf
	/*
		func Scanf(format string, a ...interface{}) (n int, err error)

	Scanf从标准输入扫描文本，根据format参数指定的格式去读取由 空白符分隔 的值保存到传递给本函数的参数中。
	本函数返回成功扫描的数据个数和遇到的任何错误。
	 */
	//fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)	//要这样输入：1:小王子 2:28 3:false
	//fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

	// 3. fmt.Scanln
	/*
		func Scanln(a ...interface{}) (n int, err error)

	Scanln类似Scan，它在遇到 换行 时才停止扫描。最后一个数据后面必须有换行或者到达结束位置。
	本函数返回成功扫描的数据个数和遇到的任何错误。
	 */
	fmt.Scanln(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
	// fmt.Scanln遇到回车就结束扫描了，这个比较常用。


	// 4. bufio.NewReader
	/*
	有时候我们想完整获取输入的内容，而输入的内容可能包含空格，这种情况下可以使用bufio包来实现。
	 */
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)

	// 5. Fscan系列
	/*
	这几个函数功能分别类似于fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，只不过它们不是从标准输入中读取数据而是从io.Reader中读取数据。
		func Fscan(r io.Reader, a ...interface{}) (n int, err error)
		func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
		func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
	 */

	// 6. Sscan系列
	/*
	这几个函数功能分别类似于fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，只不过它们不是从标准输入中读取数据而是从指定字符串中读取数据。
		func Sscan(str string, a ...interface{}) (n int, err error)
		func Sscanln(str string, a ...interface{}) (n int, err error)
		func Sscanf(str string, format string, a ...interface{}) (n int, err error)
	 */
}
