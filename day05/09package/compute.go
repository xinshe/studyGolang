package main

import (
	"fmt"

	"studyGolang/day05/09package/calc"
	pt "studyGolang/day05/09package/printTable"
)


// 包（package）
/*
包（package）是多个Go源码的集合，是一种高级的代码复用方案，Go语言为我们提供了很多内置包。

1. 定义包：
根据自己的需要创建自己的包。一个包可以简单理解为一个存放.go文件的文件夹。
该文件夹下面的所有go文件都要在代码的第一行添加如下代码，声明该文件归属的包。
	package 包名
注意事项：
	一个文件夹下面直接包含的文件只能归属一个package，同样一个package的文件不能在多个文件夹下。
	包名可以不和文件夹的名字一样，包名不能包含 - 符号。
	包名为main的包为应用程序的入口包，这种包编译后会得到一个可执行文件，而编译不包含main包的源代码则不会得到可执行文件。

2. 可见性：
如果想在一个包中引用另外一个包里的标识符（如变量、常量、类型、函数等）时，该标识符必须是对外可见的（public）。
在Go语言中只需要将标识符的首字母大写就可以让标识符对外可见了。

3. 包的导入：
要在代码中引用其他包的内容，需要使用import关键字导入使用的包。具体语法如下:
	import "包的路径"
注意事项：
	import导入语句通常放在文件开头包声明语句的下面。
	导入的包名需要使用双引号包裹起来。
	包名是从$GOPATH/src/后开始计算的，使用/进行路径分隔。
	Go语言中禁止循环导入包。

4. 自定义包名
在导入包名的时候，我们还可以为导入的包设置别名。通常用于导入的包名太长或者导入的包名冲突的情况。具体语法格式如下：
	import 别名 "包的路径"

5. 匿名导入包
如果只希望导入包，而不使用包内部的数据时，可以使用匿名导入包。具体的格式如下：
	import _ "包的路径"
匿名导入的包与其他方式导入的包一样都会被编译到可执行文件中。

6. init()初始化函数
init()函数介绍
在Go语言程序执行时导入包语句会自动触发包内部init()函数的调用。需要注意的是：
	init()函数没有参数也没有返回值。
	init()函数在程序运行时自动被调用执行，不能在代码中主动调用它。

包初始化执行的顺序如下：
	全局声明 --> init()函数 --> main()函数

init()函数执行顺序：
Go语言包会从main包开始检查其导入的所有包，每个包中又可能导入了其他的包。
Go编译器由此构建出一个树状的包引用关系，再根据引用顺序决定编译顺序，依次编译这些包的代码。

在运行时，被最后导入的包会最先初始化并调用其init()函数。
 */

func main()  {
	ret := calc.Add(10, 20)
	fmt.Println(ret)

	var stuList []pt.Student
	//stuList = make([]pt.Student, 0)
	s1 := pt.Student{
		Name:   "Andy",
		Age:    20,
		Gender: "male",
	}
	stuList = append(stuList, s1)

	s2 := pt.Student{
		Name:   "Jack",
		Age:    21,
		Gender: "female",
	}
	stuList = append(stuList, s2)

	pt.PrintStu(stuList)


}
