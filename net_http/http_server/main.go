package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(writer http.ResponseWriter, request *http.Request) {
	bytes, err := ioutil.ReadFile("net_http/http_server/page.html")
	if err != nil {
		writer.Write([]byte(fmt.Sprintf("%v", err)))
	}
	writer.Write(bytes)
}

func f2(writer http.ResponseWriter, request *http.Request) {
	// 对于Get请求，参数都是在URL上的（Query Param）
	fmt.Println(request.URL)
	queryParams := request.URL.Query()	// 解析出URL中的查询参数
	name := queryParams.Get("name")
	age := queryParams.Get("age")
	fmt.Println(name, age)

	fmt.Println(request.Method)
	fmt.Println(ioutil.ReadAll(request.Body))
	writer.Write([]byte("ok"))
}

func main()  {
	http.HandleFunc("/hello", f1)
	http.HandleFunc("/xxx", f2)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
