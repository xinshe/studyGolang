package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。

func main() {
	//var str string = "how do you do can you speak English what do you like how are you"
	var str string
	in := bufio.NewScanner(os.Stdin)
	if in.Scan() {
		str = in.Text()
	} else {
		return
	}
	splits := strings.Split(str, " ")
	var m = make(map[string]int, 3)
	for _, value := range splits {
		m[value] += 1
	}
	for k, v := range m {
		fmt.Printf("%v=%v \n", k, v)
	}
}
