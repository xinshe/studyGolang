package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Counter interface {
	Inc()
	Load() int64
}

// 普通版
type CommonCounter struct {
	value int64
}

func (c *CommonCounter) Inc()  {
	c.value++
}

func (c *CommonCounter) Load() int64 {
	return c.value
}

// 加锁版
type MutexCounter struct {
	value int64
	lock sync.Mutex
}

func (mc *MutexCounter) Inc()  {
	mc.lock.Lock()
	defer mc.lock.Unlock()
	mc.value++
}

func (mc *MutexCounter) Load() int64 {
	mc.lock.Lock()
	defer mc.lock.Unlock()
	return mc.value
}

// 原子版
type AtomicCounter struct {
	value int64
}

func (ac *AtomicCounter) Inc()  {
	atomic.AddInt64(&ac.value, 1)
}

func (ac *AtomicCounter) Load() int64 {
	return atomic.LoadInt64(&ac.value)
}

func test(c Counter) {
	start := time.Now()
	wg := sync.WaitGroup{}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}

	wg.Wait()

	end := time.Now()
	fmt.Println(end.Sub(start), c.Load())
}

func main()  {
	cc := CommonCounter{}
	test(&cc)

	mc := MutexCounter{}
	test(&mc)

	ac := AtomicCounter{}
	test(&ac)
}