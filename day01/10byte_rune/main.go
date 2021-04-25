package main

import "fmt"

/*
组成每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。 字符用单引号（’）包裹起来，如：
	var a := '中'
	var b := 'x'

Go 语言的字符有以下两种：
	uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
	rune类型，代表一个 UTF-8字符。
当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32。

Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，也可以使用 byte 型进行默认字符串处理，
性能和扩展性都有照顾。
*/

func main() {
	s := "hello复旦大学복단대학"

	for i := 0; i < len(s); i++ { // byte
		fmt.Printf("%v(%c),", s[i], s[i])
	}
	fmt.Println()

	for _, c := range s { // rune
		fmt.Printf("%v(%c),", c, c)
	}

	fmt.Println()

	c1 := "H"       // string
	c2 := 'H'       // int32
	c3 := byte('H') // uint8 (byte)
	fmt.Printf("c1:%T, c2:%T, c3:%T", c1, c2, c3)
	fmt.Println()
	fmt.Printf("c1:%v, c2:%d, c3:%d", c1, c2, c3)
	fmt.Println()

	c4 := "新" // string
	c5 := '新' // int32
	fmt.Printf("c4:%T, c5:%T", c4, c5)

}

/*
因为UTF8编码下一个中文汉字由3~4个字节组成，所以我们不能简单的按照字节去遍历一个包含中文的字符串。

字符串底层是一个byte数组，所以可以和[]byte类型相互转换。字符串是不能修改的，字符串是由byte字节组成，
所以字符串的长度是byte字节数组的长度。 rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。
*/
