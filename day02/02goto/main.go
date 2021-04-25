package main

import "fmt"

// goto(跳转到指定标签)
/*
goto语句通过标签进行代码间的无条件跳转。goto语句可以在快速跳出循环、避免重复退出上有一定的帮助。
Go语言中使用goto语句能简化一些代码的实现过程。例如双层嵌套的for循环要退出时：
*/
func main() {
	breakFlag := false
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j >= 2 {
				breakFlag = true
				break
			}
			fmt.Println(i, "\t", j)
		}
		if breakFlag {
			break
		}
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j >= 2 {
				goto BreakPos
			}
			fmt.Println(i, "\t", j)
		}
	}
BreakPos:
	fmt.Println("退出了。。。")

	/*
		break语句可以结束for、switch和select的代码块。
		break语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，
		标签要求必须定义在对应的for、switch和 select的代码块上。 举个例子：
	*/
BreakDemo:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j >= 2 {
				break BreakDemo
			}
			fmt.Println(i, "\t", j)
		}
		fmt.Println("....") // 这一句执行不到
	}

	/*
		continue语句可以结束当前循环，开始下一次的循环迭代过程，仅限在for循环内使用。
		在 continue语句后添加标签时，表示开始标签对应的循环。例如：
	*/
ContinueDemo:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j >= 2 {
				continue ContinueDemo
			}
			fmt.Println(i, "\t", j)
		}
		fmt.Println("....") // 这一句执行不到
	}

}
