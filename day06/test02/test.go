package main

import (
	"fmt"
	"runtime"
)

func TestCaller0()  {
	caller, file, line, ok := runtime.Caller(0)
	if !ok {
		return
	}
	fmt.Println(caller, file, line)
}

func TestCaller1()  {
	caller, file, line, ok := runtime.Caller(1)
	if !ok {
		return
	}
	fmt.Println(caller, file, line)
}


