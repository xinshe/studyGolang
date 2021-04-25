package main

import (
	"fmt"
	"strings"
)

// 字符串
/*
Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型（int、bool、float32、float64 等）一样。
Go语言里的字符串的内部实现使用UTF-8编码。
字符串的值为双引号(")中的内容，
可以在Go语言的源码中直接添加非ASCII码字符，例如：
s1 := "hello"
s2 := "你好"
*/

func main() {

	// 转义符
	path := "\"/Users/username/go/src\""
	fmt.Println(path)

	// 多行的字符串
	// Go语言中要定义一个多行字符串时，就必须使用 反引号 字符
	// 反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出。
	s1 := `
	世情薄
			人情恶
		雨送黄昏花易落
	`
	fmt.Println(s1)
	path2 := `"/Users/username/go/src"`
	fmt.Println(path2)

	// 字符串的操作

	// 字符串的长度
	fmt.Println(len(path2))

	// 字符串拼接
	name := "Andy"
	level := "Outstanding"
	// 方式1：+号
	ss := name + ":" + level
	fmt.Println(ss)
	// 方式2：fmt.Sprintf
	ss1 := fmt.Sprintf("%s:%s", name, level)
	fmt.Println(ss1)

	// 字符串分割
	ret := strings.Split(path2, "/")
	fmt.Println(ret)

	// 判断是否包含某个字符串
	fmt.Println(strings.Contains(ss, "And"))
	fmt.Println(strings.Contains(ss, "Ady"))

	// 前缀/后缀判断
	fmt.Println(strings.HasPrefix(ss, "And"))
	fmt.Println(strings.HasPrefix(ss, "A"))
	fmt.Println(strings.HasPrefix(ss, "Ad"))

	fmt.Println(strings.HasSuffix(ss, "ing"))
	fmt.Println(strings.HasSuffix(ss, "g"))
	fmt.Println(strings.HasSuffix(ss, "ig"))

	// 子串出现的位置
	s2 := "abcdea"
	fmt.Println(strings.Index(s2, "b"))
	fmt.Println(strings.Index(s2, "cd"))
	fmt.Println(strings.Index(s2, "acd"))
	fmt.Println(strings.LastIndex(s2, "a"))
	fmt.Println(strings.LastIndex(s2, "de"))

	// Join操作
	fmt.Println(strings.Join(ret, "++"))

}
