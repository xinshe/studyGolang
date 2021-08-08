package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁

var (
	value = 0
	lock sync.RWMutex
	wg sync.WaitGroup
)

func read()  {
	defer wg.Done()
	lock.RLock()
	fmt.Println(value)
	time.Sleep(time.Millisecond)
	lock.RUnlock()
}

func write()  {
	defer wg.Done()
	lock.Lock()
	value += 1
	time.Sleep(time.Millisecond * 2)
	lock.Unlock()
}

func main()  {
	start := time.Now()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()

	end := time.Now()
	fmt.Println(end.Sub(start))
}
