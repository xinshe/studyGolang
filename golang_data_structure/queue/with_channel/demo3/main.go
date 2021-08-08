package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var mapSync sync.Map  // key: int64, value: chan int64
var wg sync.WaitGroup
var pushPopString string
var lock sync.Mutex

func push(key, data int64) {
	defer wg.Done()
	lock.Lock()
	defer lock.Unlock()
	select {
	case getMapSyncValue(key) <- data:
		pushPopString += fmt.Sprintf("case push: key:%v value:%v\n", key, data)
	default:
		<-getMapSyncValue(key)
		pushPopString += fmt.Sprintf("pop\n")
		getMapSyncValue(key) <- data
		pushPopString += fmt.Sprintf("default push: key:%v value:%v\n", key, data)
	}
}

func getMapSyncValue(key int64) chan int64 {
	value, ok := mapSync.Load(key)
	if !ok {
		value = make(chan int64, 10)
		mapSync.Store(key, value)
	}
	return value.(chan int64)

}

func mapToString(m *sync.Map) string {
	var res string
	m.Range(func(key, value interface{}) bool {
		res += fmt.Sprintf("key:%v\tvalue:[\t",key)
		for {
			select {
			case data := <-value.(chan int64):
				res += fmt.Sprintf("%v\t", data)
			default:
				goto Flag
			}
		}
		Flag:
			res += fmt.Sprintf("]\n")
		return true
	})
	return res
}

func main() {
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		key := rand.Int63n(10)
		value := rand.Int63n(10000)
		go push(key, value)
	}
	wg.Wait()
	fmt.Println(pushPopString)
	fmt.Println(mapToString(&mapSync))
}


