package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func send(conn net.Conn) {
	defer conn.Close()	// 关闭连接
	for {
		readerStdin := bufio.NewReader(os.Stdin)
		input, err := readerStdin.ReadString('\n')
		validInput := strings.Trim(input, "\r\n")
		if strings.ToUpper(validInput) == "EXIT" {
			break
		}
		_, err = conn.Write([]byte(validInput))
		if err != nil {
			fmt.Printf("send message to client failed, err: %v", err)
			return
		}
		fmt.Printf("server: %s\n", validInput)
	}
}

func receive(conn net.Conn) {
	for {
		readerConn := bufio.NewReader(conn)		// 选择从连接中读取数据
		var buf [128]byte
		n, err := readerConn.Read(buf[:])	// 读取数据
		if err != nil {
			fmt.Printf("receive message from client failed, err: %v", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Printf("client: %s\n", recvStr)
	}
}

func main() {
	// 1. 监听端口
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("listen failed, err:%v", err)
		return
	}
	// 2. 接收客户端请求建立连接
	for {	// 不停地接收客户端的连接
		conn, err := listen.Accept()	// 建立连接
		if err != nil {
			fmt.Printf("connect failed, err:%v", err)
			continue
		}

		go receive(conn)	// 启动一个goroutine进行接收
		go send(conn)
	}
}
