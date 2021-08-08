package main

import (
	"fmt"
	"strconv"
	"sync"
)
// 并发不安全的map

var m = make(map[string]int)

func set(key string, value int)  {
	m[key] = value
}

func get(key string) int {
	return m[key]
}

func main()  {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := strconv.Itoa(n)
			value := n * n
			set(key, value)
			fmt.Printf("key:%v value:%v\n", key, get(key))
		}(i)
	}
	wg.Wait()

}
