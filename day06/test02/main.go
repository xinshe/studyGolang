package main

import (
	"fmt"
	"runtime"
)

func main()  {
	f1()
	f1()
	TestCaller0()
	TestCaller1()
}

func f1()  {
	caller, file, line, ok := runtime.Caller(1)
	if !ok {
		return
	}
	fmt.Println(caller, file, line)
}
