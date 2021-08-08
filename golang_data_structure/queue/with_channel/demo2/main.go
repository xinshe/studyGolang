package main

import (
	"math/rand"
	"sync"
	"sync/atomic"
)

var mapSync atomic.Value
var wg sync.WaitGroup
var lock sync.Mutex

func push(key, value int64) {
	defer wg.Done()
	lock.Lock()
	defer lock.Unlock()
	if _, ok := getMapSync()[key]; !ok {
		getMapSync()[key] = make(chan int64, 10)
	}

	//if len(getMapSync()[key]) >= 10 {
	//	<-getMapSync()[key]
	//}
	//getMapSync()[key] <- value

	select {
	case getMapSync()[key] <- value:
	default:
		<-getMapSync()[key]
		getMapSync()[key] <- value
	}
}

func getMapSync() map[int64]chan int64 {
	return mapSync.Load().(map[int64]chan int64)
}

func main()  {
	mapSync.Store(map[int64]chan int64{})
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		key := rand.Int63n(10000)
		value := rand.Int63n(10000)
		go func() {
			push(key, value)
		}()
	}
	wg.Wait()
}
