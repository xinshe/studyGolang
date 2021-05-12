package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

/*
文件写入操作

os.OpenFile()函数能够以指定模式打开文件，从而实现文件写入相关功能。
	func OpenFile(name string, flag int, perm FileMode) (*File, error) {
		...
	}
其中：
name：要打开的文件名 flag：打开文件的模式。模式有以下几种：
	模式			含义
	os.O_WRONLY	只写
	os.O_CREATE	创建文件
	os.O_RDONLY	只读
	os.O_RDWR	读写
	os.O_TRUNC	清空
	os.O_APPEND	追加
perm：文件权限，一个八进制数。r（读）04，w（写）02，x（执行）01。
 */

// Write和WriteString
func writeDemo1()  {
	file, err := os.OpenFile("day05/11file_write/a.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	content := "hello world! 中国\n"
	file.Write([]byte(content))	// 写入字节切片数据
	file.WriteString(content)	// 直接写入字符串数据
}

// bufio.NewWriter
func writeDemo2()  {
	file, err := os.OpenFile("day05/11file_write/a.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	content := "I love you\n"
	for i := 0; i < 10; i++ {
		writer.Write([]byte(content))	// 将数据先写入缓存
	}
	writer.Flush()	// 将缓存中的内容写入文件
}

// ioutil.WriteFile
func writeDemo3()  {
	content := "All is well\n"
	err := ioutil.WriteFile("day05/11file_write/a.txt", []byte(content), 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
}

func main()  {
	//writeDemo1()
	//writeDemo2()
	writeDemo3()
}

