package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func fileInsert(filePath string, index int64, content string) {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}

	tmpPath := getTmpFilePath(filePath)
	tmpFile, err := os.OpenFile(tmpPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open tmp file failed, err:%v\n", err)
		return
	}

	var buf [1]byte
	var curPos = 0
	for int64(curPos) < index {
		n, err := file.Read(buf[:])
		curPos += n
		if err == io.EOF {
			tmpFile.Write(buf[:n])
			fmt.Printf("1. copy file done, err:%v\n", err)
			break
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%v\n", err)
			return
		}
		tmpFile.Write(buf[:n])
	}

	// 插入数据
	tmpFile.Write([]byte(content))
	fmt.Println("2. insert file done")

	// 从index位置开始读取file，并写入tmp文件
	file.Seek(index, 0) // 光标移动到索引为index的位置
	var buf2 [1024]byte
	for {
		n, err := file.Read(buf2[:])
		if err == io.EOF {
			tmpFile.Write(buf2[:n])
			fmt.Printf("3. copy file done. err:%v\n", err)
			break
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%v\n", err)
			return
		}
		tmpFile.Write(buf2[:n])
	}

	tmpFile.Close()
	file.Close()

	err = os.Rename(tmpPath, filePath)
	if err != nil {
		fmt.Printf("rename file failed, err:%v", err)
	}
	fmt.Println("finished!")
	return
}

func getTmpFilePath(filePath string) string  {
	split := strings.Split(filePath, "/")
	var tmpPath string
	var fileDir string
	fileName := split[len(split) - 1]
	for i := 0; i < len(split) - 1; i++ {
		fileDir += split[i]
		fileDir += "/"
	}
	tmpPath = fileDir + fileName + ".tmp"
	return tmpPath
}

func main() {
	filePath := "day05/13file_insert/a.txt"
	content := " xxxxx "
	fileInsert(filePath, 100, content)
}
