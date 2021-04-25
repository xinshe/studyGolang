package main

import "fmt"

// 整型
/*
整型分为以下两个大类： 按长度分为：int8、int16、int32、int64 对应的无符号整型：uint8、uint16、uint32、uint64
其中，uint8就是我们熟知的byte型，int16对应C语言中的short型，int64对应C语言中的long型。

类型		描述
uint8	无符号 8位整型 (0 到 255)
uint16	无符号 16位整型 (0 到 65535)
uint32	无符号 32位整型 (0 到 4294967295)
uint64	无符号 64位整型 (0 到 18446744073709551615)
int8	有符号 8位整型 (-128 到 127)
int16	有符号 16位整型 (-32768 到 32767)
int32	有符号 32位整型 (-2147483648 到 2147483647)
int64	有符号 64位整型 (-9223372036854775808 到 9223372036854775807)
*/

// 特殊整型
/*
类型		描述
uint	32位操作系统上就是uint32，64位操作系统上就是uint64
int		32位操作系统上就是int32，64位操作系统上就是int64
uintptr	无符号整型，用于存放一个指针
注意： 在使用int和uint类型时，不能假定它是32位或64位的整型，而是考虑int和uint可能在不同平台上的差异。

注意事项 获取对象的长度的内建len()函数返回的长度可以根据不同平台的字节长度进行变化。实际使用中，切片或 map 的元素数量等
都可以用int来表示。在涉及到二进制传输、读写文件的结构描述时，为了保持文件的结构不会受到不同编译目标平台字节长度的影响，
不要使用int和uint。
*/

func main() {
	// 十进制
	i1 := 102
	fmt.Printf("%d \n", i1)
	fmt.Printf("%b \n", i1) // 十进制转二进制
	fmt.Printf("%o \n", i1) // 十进制转八进制
	fmt.Printf("%x \n", i1) // 十进制转十六进制

	// 八进制，以0开头
	i2 := 077
	fmt.Printf("%d \n", i2) // 八进制转十进制
	fmt.Printf("%b \n", i2) // 八进制转二进制
	fmt.Printf("%o \n", i2)
	fmt.Printf("%x \n", i2) // 八进制转十六进制

	// 十六进制，以0x开头
	//var i3 int = 0xff
	i3 := 0xff
	fmt.Printf("%d \n", i3) // 十六进制转十进制
	fmt.Printf("%b \n", i3) // 十六进制转二进制
	fmt.Printf("%o \n", i3) // 十六进制转八进制
	fmt.Printf("%x \n", i3)
	fmt.Printf("%X \n", i3)

	// 查看变量类型
	fmt.Printf("%T \n", i1)
	fmt.Printf("%T \n", i2)
	fmt.Printf("%T \n", i3)

	// 声明一个int8类型的变量
	i4 := int8(125)
	fmt.Printf("%d \n", i4)
	fmt.Printf("%T \n", i4)

	var i5 int8 = 76
	fmt.Printf("%d \n", i5)
	fmt.Printf("%T \n", i5)

}
