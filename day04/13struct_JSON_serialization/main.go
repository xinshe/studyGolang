package main

import (
	"encoding/json"
	"fmt"
)

// 结构体 和 JSON 互转
/*
序列化：结构体 --> JSON
反序列化：JSON --> 结构体
 */
/*
结构体标签（Tag）结构体的元信息，可以在运行的时候通过反射的机制读取出来。 Tag在结构体字段的后方定义，
由一对反引号包裹起来，具体的格式如下：
	`key1:"value1" key2:"value2"`

结构体tag由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。同一个结构体字段可以设置多个键值对tag，
不同的键值对之间使用空格分隔。

注意事项： 为结构体编写Tag时，必须严格遵守键值对的规则。结构体标签的解析代码的容错能力很差，一旦格式写错，
编译和运行时都不会提示任何错误，通过反射也无法正确取值。例如不要在key和value之间添加空格。
 */

// 字段名首字母大写表示包外可访问，即会参与到序列化中
type person struct {
	Name string 	`json:"name"`
	Age int			`json:"age"`
	gender string
}

func main()  {
	p1 := person{
		Name: "Andy",
		Age:  19,
		gender: "男",
	}
	// 序列化
	str, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err:%v", err)
	}
	fmt.Println(string(str))	// {"name":"Andy","age":19}

	// 反序列化
	s := `{"name":"Jack", "age":20, "gender":"male"}`
	var p2 person
	json.Unmarshal([]byte(s), &p2)//这里传指针
	fmt.Printf("%#v\n", p2)	// main.person{Name:"Jack", Age:20, gender:""}
}
