package main

import "fmt"

// 观察下面代码，写出最终的打印结果。

func main() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s) // [1 2 3]
	m["q1mi"] = s
	fmt.Println(len(m["q1mi"]), cap(m["q1mi"])) // 3 4
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)         // [1 3]
	fmt.Printf("%+v\n", m["q1mi"]) // [1 3 3] ???
}
