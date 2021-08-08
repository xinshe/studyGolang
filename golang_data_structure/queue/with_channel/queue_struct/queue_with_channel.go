package main

import (
	"fmt"
	"time"
)

type Queue struct {
	ch chan int64
	size int64
}

func InitQueue(size int64) *Queue {
	return &Queue{
		ch:   make(chan int64, size),
		size: size,
	}
}

func (q Queue) push(data int64) {
	select {
	case q.ch <- data:
		fmt.Printf("Push %v to queue successful.\n", data)
	default:
		pollData := <-q.ch
		fmt.Printf("Queue is full, and then pop %v from queue successful.\n", pollData)
		q.ch <- data
		fmt.Printf("Repush %v to queue successful.\n", data)
	}
}

func (q Queue) pop() {
	select {
	case pollData := <-q.ch:
		fmt.Printf("Pop %v from queue successful.\n", pollData)
	default:
		fmt.Printf("Queue is empty, pop from queue failed.\n")
	}
}

func main()  {
	queue := InitQueue(2)

	/*
	queue.push(1)
	queue.push(2)
	queue.push(3)
	queue.push(4)
	queue.pop()
	queue.pop()
	queue.pop()
	queue.pop()
	*/

	///*
	for i := 0; i < 10; i++ {
		go func() {
			queue.push(1)
			queue.push(2)
		}()
	}
		time.Sleep(time.Second * 5)
	 //*/

	/*
	var queueSync atomic.Value
	queueSync.Store(queue)
	for i := 0; i < 100; i++ {
		go func() {
			queueSync.Load().(*Queue).push(1)
			queueSync.Load().(*Queue).push(2)
		}()

	}
	time.Sleep(time.Second * 10)
	*/

}