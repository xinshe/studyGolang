package main

import "fmt"

func main() {
	// 元素为map类型的切片
	var s = make([]map[string]int, 5, 10)
	s[0] = map[string]int{"shanghai": 98, "beijing": 97}
	s[0]["shenzhen"] = 96
	s[1] = make(map[string]int, 12)
	s[1]["hangzhou"] = 86
	s[1]["chengdu"] = 88
	fmt.Println(s) // [map[beijing:97 shanghai:98 shenzhen:96] map[chengdu:88 hangzhou:86] map[] map[] map[]]

	// 值为切片类型的map
	var m = make(map[string][]int, 10)
	m["西瓜"] = []int{2, 4, 6}
	m["葡萄"] = make([]int, 3, 8)
	fmt.Println(m) // map[葡萄:[0 0 0] 西瓜:[2 4 6]]
	m["西瓜"][0] = 100
	m["葡萄"][1] = 456
	fmt.Println(m) // map[葡萄:[0 456 0] 西瓜:[100 4 6]]
}
