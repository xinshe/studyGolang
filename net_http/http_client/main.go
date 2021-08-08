package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {
	resp, err := http.Get("http://127.0.0.1:9090/xxx?name=Andy&age=19")	// 不要忘了加 http://
	if err != nil {
		fmt.Printf("get url failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	// 从resp中把服务端返回的数据读出来
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(bytes))	// 在控制台打印从resp中返回的数据
}
