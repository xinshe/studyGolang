package main

import "fmt"

func main() {
	arr := []int{30000, 24000, 3000, 2800, 2000,
		1600, 2000, 18000}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Println(sum)
}