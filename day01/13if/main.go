package main

import "fmt"

// if 判断

func main() {
	age := 20
	if age >= 18 {
		fmt.Println("成年了")
	} else {
		fmt.Println("未成年")
	}

	if age >= 40 {
		fmt.Println("人到中年")
	} else if age >= 18 {
		fmt.Println("成年了")
	} else {
		fmt.Println("未成年")
	}

	// score作用域仅在该判断块中，这样做可以节省内存开销
	if score := 89; score >= 60 {
		fmt.Println("及格了")
	} else {
		fmt.Println("没有及格")
	}

}
