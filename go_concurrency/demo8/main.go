package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
使用goroutine和channel实现一个计算int64随机数各位数和的程序。
	开启一个goroutine循环生成int64类型的随机数，发送到jobChan
	开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
	主goroutine从resultChan取出结果并打印到终端输出
 */

type job struct {
	value int64
}

type result struct {
	job
	sum int64
}

func produce(jobs chan<- *job) {
	for {
		rand.Seed(time.Now().UnixNano())
		r := rand.Int63()
		jobs <- &job{
			value: r,
		}
		time.Sleep(time.Second)
	}
}

func worker(jobs <-chan *job, results chan<- *result) {
	for jc := range jobs {
		n := jc.value
		sum := int64(0)
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		results <- &result{
			job: *jc,
			sum: sum,
		}
	}
}

func main()  {

	var jobChan = make(chan *job, 32)
	var resultChan = make(chan *result, 32)

	go produce(jobChan)

	for i := 0; i < 24; i++ {
		go worker(jobChan, resultChan)
	}

	for rc := range resultChan {
		fmt.Printf("value:%d sum:%d\n", rc.value, rc.sum)
	}
}
