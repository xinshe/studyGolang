package main

import (
	"time"
)

func main()  {
	// 1. 获取当前时间，格式化输出为2017/06/19 20:30:05`格式。
	println(time.Now().Format("2006/01/02 15:04:05"))

}
