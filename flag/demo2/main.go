package main

import (
	"flag"
	"fmt"
	"time"
)

func main()  {
	var name string
	var age int
	var married bool
	var delay time.Duration

	flag.StringVar(&name, "name", "Andy", "姓名")
	flag.IntVar(&age, "age", 24, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "delay", 10000, "时长")

	// 解析命令行参数
	flag.Parse()

	fmt.Println(name, age, married, delay)

	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())



}
