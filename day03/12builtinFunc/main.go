package main

import "fmt"

/*
内置函数
close			主要用来关闭channel
len				用来求长度，比如string、array、slice、map、channel
new				用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针
make			用来分配内存，主要用来分配引用类型，比如chan、map、slice
append			用来追加元素到数组、slice中
panic和recover	用来做错误处理
*/
func main() {
	var p = new(int)
	*p++
	fmt.Println(*p)

	var m = make(map[string]string, 10)
	m["hello"] = "world"
	fmt.Println(m)
}
