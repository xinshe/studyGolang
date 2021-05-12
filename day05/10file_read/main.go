package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 文件的打开、关闭、读操作
/*
计算机中的文件是存储在外部介质（通常是磁盘）上的数据集合，文件分为文本文件和二进制文件。
 */

// 读取文件方式1：file.Read()
/*
Read方法 接收一个字节切片，返回读取的字节数和可能的具体错误，读到文件末尾时会返回0和io.EOF
 */
func readFile() {
	file, err := os.Open("day05/10file_read/a.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer file.Close()

	var content []byte
	tmp := make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if  err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}

// 读取文件方式2：bufio读取文件
/*
bufio是在file的基础上封装了一层API，支持更多的功能。
 */
func readFileByBufio()  {
	file, err := os.Open("day05/10file_read/a.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return	
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for  {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Print(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}
}

// 读取文件方式3：ioutil读取整个文件
func readFileByIoutil()  {
	fileContent, err := ioutil.ReadFile("day05/10file_read/a.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	fmt.Print(string(fileContent))

}

func main()  {
	// 只读方式打开当前目录下的main.go文件
	file, err := os.Open("day05/10file_read/main.go")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	// 关闭文件
	defer file.Close()

	//readFile()

	//readFileByBufio()

	readFileByIoutil()
}
