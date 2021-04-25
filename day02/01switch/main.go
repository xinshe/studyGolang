package main

import "fmt"

// switch case
/*
使用switch语句可方便地对大量的值进行条件判断。

*/
func main() {
	finger := 2
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("输入无效！")
	}
	// Go语言规定每个switch只能有一个default分支。

	// 一个分支可以有多个值，多个case值中间使用英文逗号分隔。
	switch n := 4; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("odd")
	case 0, 2, 4, 6, 8:
		fmt.Println("even")
	default:
		fmt.Println("amazing!")
	}

	// 分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量。例如：
	age := 24
	switch {
	case age <= 25:
		fmt.Println("day day study")
	case age > 25 && age <= 50:
		fmt.Println("day day work")
	case age > 50:
		fmt.Println("day day happy")
	default:
		fmt.Println("day day up")
	}

	// fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。
	s := "a"
	switch s {
	case "a":
		fmt.Println("a")
		fallthrough
	case "b":
		fmt.Println("b")
	case "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}

	// goto语句通过标签进行代码间的无条件跳转。goto语句可以在快速跳出循环、避免重复退出上有一定的帮助。
	// Go语言中使用goto语句能简化一些代码的实现过程。 例如双层嵌套的for循环要退出时：
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				breakFlag = true
				break
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		// 外层for循环判断
		if breakFlag {
			break
		}
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	// 标签
breakTag:
	fmt.Println("结束for循环")
}
