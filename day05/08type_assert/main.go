package main

import "fmt"

// 类型断言
/*
空接口可以存储任意类型的值，那我们如何获取其存储的具体数据呢？

接口值:
一个接口的值（简称接口值）是由一个 具体类型 和 具体类型的值 两部分组成的。这两部分分别称为接口的 动态类型 和 动态值。

想要判断空接口中的值这个时候就可以使用类型断言，其语法格式：
	x.(T)
其中：
	x：表示类型为interface{}的变量
	T：表示断言x可能是的类型。
该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败。

因为空接口可以存储任意类型值的特点，所以空接口在Go语言中的使用十分广泛。

关于接口需要注意的是，只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。
不要为了接口而写接口，那样只会增加不必要的抽象，导致不必要的运行时损耗。

*/

func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupported type！")
	}
}

func main() {
	var x interface{}
	x = "hello"
	v, ok := x.(string)
	if !ok {
		fmt.Println("failure")
	} else {
		fmt.Printf("%T, %v\n", v, v)
	}
}
