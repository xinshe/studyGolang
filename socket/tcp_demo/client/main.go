package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func send(conn net.Conn) {
	defer conn.Close()	// 关闭连接
	defer wg.Done()
	for {
		readerStdin := bufio.NewReader(os.Stdin)	// 读取用户输入
		input, err := readerStdin.ReadString('\n')
		validInput := strings.Trim(input, "\r\n")
		if strings.ToUpper(validInput) == "EXIT" {
			break
		}
		_, err = conn.Write([]byte(validInput))		// 发送数据
		if err != nil {
			fmt.Printf("send message to server failed, err: %v", err)
			break
		}
		fmt.Printf("client: %s\n", validInput)
	}
}

func receive(conn net.Conn) {
	defer wg.Done()
	for {
		readerConn := bufio.NewReader(conn)		// 选择从连接中读取数据
		var buf [128]byte
		n, err := readerConn.Read(buf[:])	// 读取数据
		if err != nil {
			fmt.Printf("receive message from server failed, err: %v", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Printf("server: %s\n", recvStr)
	}
}

func main()  {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("client connect failed, err: %v", err)
		return
	}

	wg.Add(2)
	go send(conn)
	go receive(conn)
	wg.Wait()
}
