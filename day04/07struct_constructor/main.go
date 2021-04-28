package main

// 构造函数
/*
golang中要自己定义和实现构造函数

构造函数：约定成俗用 new 开头
返回的是结构体还是结构体指针: 当结构体比较大的时候尽量使用结构体指针，减少程序的內存开销
 */

type person struct {
	name string
	age int
	gender string
}

type dog struct {
	name string
}

func newPerson(name string, age int, gender string) *person {
	return &person{
		name:   name,
		age:    age,
		gender: gender,
	}
}

func newDog(name string) dog {
	return dog{name: name}
}

func main()  {

}
