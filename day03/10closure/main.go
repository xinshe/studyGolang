package main

import "fmt"

func f1(f func()) {
	fmt.Println("f1 done")
	f()
}

func f2(x, y int) {
	fmt.Println("f2 done")
	fmt.Println(x + y)
}

// 要求实现这样的调用：f1(f2)
func f3(f func(int, int), x, y int) func() {
	fmt.Println("f3 done")
	return func() {
		f(x, y)
	}
}

func main() {
	f1(f3(f2, 100, 200))
	/*
		f3 done
		f1 done
		f2 done
		300
	*/
}
