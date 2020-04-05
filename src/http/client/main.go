package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get、Head、Post和PostForm函数发出HTTP/HTTPS请求。
func main() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("get url:https://www.baidu.com response error", err)
		return
	}
	defer resp.Body.Close()

	// 从服务器把响应数据读取出来
	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("read url error", err)
		return
	}
	fmt.Println(string(b))
}
