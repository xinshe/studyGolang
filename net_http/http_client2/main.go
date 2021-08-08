package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main()  {
	data := url.Values{}
	data.Set("name", "John")
	data.Set("age", "23")
	queryStr := data.Encode()
	fmt.Println(queryStr)

	urlObj, _ := url.Parse("http://127.0.0.1:9090/xxx")
	urlObj.RawQuery = queryStr

	request, _ := http.NewRequest("GET", urlObj.String(), nil)

	resp, _ := http.DefaultClient.Do(request)

	defer resp.Body.Close()
	// 从resp中把服务端返回的数据读出来
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(bytes))	// 在控制台打印从resp中返回的数据
}
