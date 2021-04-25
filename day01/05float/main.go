package main

import (
	"fmt"
	"math"
)

// 浮点型

func main() {
	f1 := 1.23856
	fmt.Printf("%f \n", f1)   // 1.23856
	fmt.Printf("%.2f \n", f1) // 1.24
	fmt.Printf("%T \n", f1)   // float64 浮点数默认声明为float64类型

	// 显式声明为float32
	f2 := float32(3.14159)
	fmt.Printf("%f \n", f2)   // 3.141590
	fmt.Printf("%.2f \n", f2) // 3.14
	fmt.Printf("%T \n", f2)   // float32

	// 浮点数的最大值
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

}
