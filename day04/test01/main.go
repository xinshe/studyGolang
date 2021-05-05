package main

import (
	"fmt"
	"os"
)

/*
使用“面向对象”的思维方式编写一个学生信息管理系统。
	学生有id、姓名、年龄、分数等信息
	程序提供展示学生列表、添加学生、编辑学生信息、删除学生等功能
 */


var stuList map[int]*student

type student struct {
	id int
	name string
	age int
	score float64
}

func showStuList()  {
	fmt.Println("id\tname\tage\tscore")
	for k, v := range stuList {
		fmt.Printf("%d\t%s\t%d\t%f\n", k, v.name, v.age, v.score)
	}
}

func addStu(stu student) bool {
	_, ok := stuList[stu.id]
	if ok {
		fmt.Println("该学生id已存在，请重新输入")
		return false
	} else {
		stuList[stu.id] = &stu
		return true
	}
}

func editStuScore(id int, score float64) bool {
	_, ok := stuList[id]
	if ok {
		stuList[id].score = score
		return true
	} else {
		fmt.Println("该学生id不存在，请重新输入")
		return false
	}
}

func removeStu(id int) bool {
	_, ok := stuList[id]
	if ok {
		delete(stuList, id)
		return true
	} else {
		fmt.Println("该学生id不存在，请重新输入")
		return false
	}
}

func main()  {
	stuList = make(map[int]*student, 100)
	for {
		fmt.Println("欢迎进入学生信息管理系统")
		fmt.Println(`请选择操作：
1. 展示学生列表
2. 添加学生
3. 编辑学生信息
4. 删除学生
5. 退出系统`)
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Printf("你选择的操作是 %d.展示学生列表\n", choice)
			showStuList()
		case 2:
			fmt.Printf("你选择的操作是 %d.添加学生\n", choice)
			var id, age int
			var name string
			var score float64
			fmt.Println("请输入学生id")
			fmt.Scanln(&id)
			fmt.Println("请输入学生姓名")
			fmt.Scanln(&name)
			fmt.Println("请输入学生年龄")
			fmt.Scanln(&age)
			fmt.Println("请输入学生分数")
			fmt.Scanln(&score)
			s := student{
				id:    id,
				name:  name,
				age:   age,
				score: score,
			}
			addStu(s)
		case 3:
			fmt.Printf("你选择的操作是 %d.编辑学生信息\n", choice)
			var id int
			var score float64
			fmt.Println("请输入学生id")
			fmt.Scanln(&id)
			fmt.Println("请输入学生分数")
			fmt.Scanln(&score)
			editStuScore(id, score)
		case 4:
			fmt.Printf("你选择的操作是 %d.删除学生\n", choice)
			var id int
			fmt.Println("请输入学生id")
			fmt.Scanln(&id)
			removeStu(id)
		case 5:
			fmt.Printf("你选择的操作是 %d.退出系统\n", choice)
			os.Exit(1)
		default:
			fmt.Printf("输入的操作 %d 为无效操作\n", choice)
		}
		fmt.Println("-------------------------------------")
	}
}

