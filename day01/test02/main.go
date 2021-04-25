package main

import (
	"fmt"
	"unicode"
)

/*
编写代码统计出字符串"hello沙河小王子복단대학"中汉字的数量。
*/
func main() {
	s := "hello沙河小王子복단대학"
	count := 0
	for _, c := range s {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	fmt.Println(count)
}
