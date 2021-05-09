package printTable

import "fmt"

type Student struct {	// 要想在包外可以访问，结构体名要大写，字段亦如此
	Name string
	Age int
	Gender string
}

func PrintStu(stuList []Student) {
	fmt.Printf("%s\t%s\t%s\n", "Name", "Age", "Gender")
	for _, v := range stuList {
		fmt.Printf("%s\t%d\t%s\n", v.Name, v.Age, v.Gender)
	}

}
