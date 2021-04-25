package main

import "fmt"

func main() {
	var s = make([]string, 5, 10)
	var s1 = make([]int, 5, 10)
	for i := 1; i < 10; i++ {
		s = append(s, fmt.Sprintf("%v", i))
		s1 = append(s1, i)
	}
	fmt.Println(s)  // [     1 2 3 4 5 6 7 8 9]
	fmt.Println(s1) // [0 0 0 0 0 1 2 3 4 5 6 7 8 9
}
