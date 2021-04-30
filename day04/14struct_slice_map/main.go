package main

import "fmt"

/*
因为slice和map这两种数据类型都包含了指向底层数据的指针，因此我们在需要复制它们时要特别注意。我们来看下面的例子：
 */
type person struct {
	name string
	age int
	hobby []string
}

// 不正确的拷贝方式
//func (p *person) setHobby(hobby []string)  {
//	p.hobby = hobby
//}

// 正确的做法是在方法中使用传入的slice的拷贝进行结构体赋值。
func (p *person) setHobby(hobby []string)  {
	p.hobby = make([]string, len(hobby))
	copy(p.hobby, hobby)
}

// 同样的问题也存在于返回值slice和map的情况，在实际编码过程中一定要注意这个问题。
func main()  {
	p1 := person{name: "Andy", age: 19}
	data := []string{"狼人杀", "羽毛球", "手游"}
	p1.setHobby(data)
	fmt.Println(p1.hobby)

	data[1] = "篮球"
	fmt.Println(p1.hobby)

}
