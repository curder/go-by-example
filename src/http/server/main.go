package main

import (
	"fmt"
	"net/http"
)

// ListenAndServe使用指定的监听地址和处理器启动一个HTTP服务端。
// 处理器参数通常是nil，这表示采用包变量DefaultServeMux作为处理器。

// 1. http.HandleFunc 设置处理的请求路径的回调函数，在回调函数中通过 http.ResponseWrite 处理响应
// 2. http.ListenAndServe 开启服务器和端口监听服务
// 3. 使用浏览器请求对应IP、端口和请求路径，对http发送请求
func main() {
	// Handle和HandleFunc函数可以向DefaultServeMux添加处理器
	http.HandleFunc("/", indexHandle) // 设置访问路径的处理函数

	fmt.Println("http://127.0.0.1:9090")
	err := http.ListenAndServe("127.0.0.1:9090", nil) // 配置本地服务器的IP地址和监听端口


	if err != nil {
		fmt.Println("listen server 127.0.0.1:9090 error", err)
		return
	}
}

// 处理请求的函数
func indexHandle(w http.ResponseWriter, r *http.Request)  {
	str := `<h1 style="color: gray">hello world</h1>` // 准备要返回的字符串

	_, err := w.Write([]byte(str)) // 返回切片

	if err != nil {
		fmt.Println("indexHandle error: ", err)
	}
}
