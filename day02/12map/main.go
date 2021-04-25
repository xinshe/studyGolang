package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// map的基本使用
	var m1 map[string]int
	fmt.Println(m1 == nil)        // true	(还没有初始化，没有在内存中开辟空间)
	m1 = make(map[string]int, 10) // 要估算好该map容量，避免在程序运行期间再动态扩容
	m1["苹果"] = 19
	m1["西瓜"] = 27
	fmt.Println(m1) // map[苹果:19 西瓜:27]

	fmt.Println(m1["西瓜"])
	fmt.Println(m1["草莓"]) // 如果不存在这个key，拿到对应值类型的零值

	// 判断某个键是否存在
	value, ok := m1["香蕉"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("查无此key")
	}

	// map的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	for k, _ := range m1 {
		fmt.Println(k)
	}

	for _, v := range m1 {
		fmt.Println(v)
	}

	// 使用delete()函数删除键值对，内建函数delete按照指定的键将元素从映射中删除。若m为nil或无此元素，delete不进行操作。
	delete(m1, "西瓜")
	fmt.Println(m1)
	delete(m1, "火龙果")

	// 按照指定顺序遍历map
	rand.Seed(time.Now().UnixNano())
	var scoreMap = make(map[string]int, 100)
	for i := 0; i < 100; i++ {
		stu := fmt.Sprintf("stu00%v", i)
		scoreMap[stu] = rand.Intn(100)
	}

	var keys = make([]string, 0, 100) // 长度要记得设为0
	for k, _ := range scoreMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, v := range keys {
		fmt.Println(v, scoreMap[v])
	}
}
