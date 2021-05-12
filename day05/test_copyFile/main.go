package main

import (
	"fmt"
	"io"
	"os"
)

// 借助io.Copy()实现一个拷贝文件函数。

func copyFile(dstPath, srcPath string) (written int64, err error) {
	// 以读方式打开源文件
	srcFile, err := os.Open(srcPath)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", srcPath, err)
		return
	}
	defer srcFile.Close()

	// 以写|创建的方式打开目标文件
	dstFile, err := os.OpenFile(dstPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", dstPath, err)
		return
	}
	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)	// //调用io.Copy()拷贝内容
}

func main()  {
	srcPath := "day05/test_copyFile/src.txt"
	dstPath := "day05/test_copyFile/dst.txt"
	n, err := copyFile(dstPath, srcPath)
	if err != nil {
		fmt.Println("copy file failed, err:", err)
		return
	}
	fmt.Println("copy ", n)
}
