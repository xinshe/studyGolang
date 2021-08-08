package main

import (
	"bufio"
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"strings"
)

// NSQ Producer Demo

var producer *nsq.Producer

// 初始化生产者
func initConfig(addr string) (err error) {
	config := nsq.NewConfig()
	producer, err = nsq.NewProducer(addr, config)
	if err != nil {
		fmt.Printf("new producer failed, err:%v\n", err)
		return err
	}
	return nil
}

func main() {
	addr := "127.0.0.1:4150"
	err := initConfig(addr)
	if err != nil {
		fmt.Printf("init config failed, err:%v\n", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)	// 从标准输入读取

	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read string from os.Stdin failed, err:%v\n", err)
			continue
		}
		data = strings.TrimSpace(data)
		if data == "Q" {	// 输入Q退出
			break
		}
		topic := fmt.Sprintf("topic-%v", os.Getpid())
		err = producer.Publish(topic, []byte(data))
		if err != nil {
			fmt.Printf("publish %v to %v failed. err:%v\n", data, topic, err)
			continue
		}
	}
}
