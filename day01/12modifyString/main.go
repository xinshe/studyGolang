package main

import "fmt"

/*
要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
*/
func main() {
	s1 := "红萝卜"
	s2 := []rune(s1)
	s2[0] = '白'
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(string(s2))

	s3 := "hello"
	s4 := []byte(s3)
	s4[3] = 'm'
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(string(s4))

}
