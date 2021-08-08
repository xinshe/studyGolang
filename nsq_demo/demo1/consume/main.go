package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// NSQ Consumer Demo

// MyHandler 是一个消费者类型
type MyHandler struct {
	title string
}

// HandleMessage 是需要实现的处理消息的方法
func (h MyHandler) HandleMessage(msg *nsq.Message) (err error) {
	fmt.Printf("%v recvive message from %v, message:%v \n", h.title, msg.NSQDAddress, string(msg.Body))
	return
}

// 初始化消费者
func initConsumer(topic, channel, addr string) (err error) {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		fmt.Printf("create consumer failed, err:%v\n", err)
		return err
	}

	handler := &MyHandler{
		title: "消费者1",
	}

	consumer.AddHandler(handler)

	// if err := c.ConnectToNSQD(address); err != nil { // 直接连 NSQD
	if err := consumer.ConnectToNSQLookupd(addr); err != nil {	// 通过lookupd查询
		return err
	}
	return nil
}

func main()  {
	err := initConsumer("topic-48115", "first-channel", "127.0.0.1:4161")
	if err != nil {
		fmt.Printf("init consumer failed, err:%v\n", err)
		return
	}
	c := make(chan os.Signal)        // 定义一个信号的通道
	signal.Notify(c, syscall.SIGINT) // 转发键盘中断信号到c
	<-c                              // 阻塞
}