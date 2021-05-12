package main

import (
	"bufio"
	"fmt"
	"os"
)

// 从控制台获取输入

// fmt.Scanln方式
func inputDemo1()  {
	fmt.Println("请输入...")
	var s string
	fmt.Scanln(&s)	// 获取到换行符就暂停输入，但是输入中有空格也会暂停输入，也就是说无法读入空格
	fmt.Println(s)
}

// bufio.NewReader方式
func inputDemo2()  {
	fmt.Println("请输入...")
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')	// 读取到换行符就暂停，可以读入空格
	fmt.Println(s)

}

func main()  {
	//inputDemo1()
	inputDemo2()
}

