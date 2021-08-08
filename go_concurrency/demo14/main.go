package main

import (
	"fmt"
	"strconv"
	"sync"
)

// sync.Map

func main()  {
	m := sync.Map{}
	wg := sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := strconv.Itoa(n)
			m.Store(key, n * 3)
			value, _ := m.Load(key)
			fmt.Printf("key:%v value:%v\n", key, value)
		}(i)
	}
	wg.Wait()
	
}
