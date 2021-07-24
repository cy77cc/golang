package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HandleWhoami(responseWriter http.ResponseWriter, request *http.Request) {
	requestMssageMap := make(map[string]interface{})
	requestMssageMap["客户端IP"] = request.RemoteAddr
	requestMssageMap["协议"] = request.Proto
	requestMssageMap["请求方法"] = request.Method
	requestMssageMap["请求的主机"] = request.Host
	requestMssageMap["请求的URL"] = request.URL.String()
	requestMssageMap["数据长度"] = request.ContentLength
	marshal, err := json.Marshal(requestMssageMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, _ = responseWriter.Write(marshal)

}

func main() {
	//pattern是路由
	http.HandleFunc("/hello", func(responseWriter http.ResponseWriter, request *http.Request) {
		_, _ = responseWriter.Write([]byte("Hello"))
	})

	http.HandleFunc("/whoami", HandleWhoami)

	http.HandleFunc("/hupu", func(responseWriter http.ResponseWriter, request *http.Request) {
		header := responseWriter.Header()
		header.Add("Content-Type", "text/html;charset=UTF-8")
		bytes, _ := ioutil.ReadFile(`E:\golang\网络编程\HTTP编程\hupu.html`)
		_, _ = responseWriter.Write(bytes)
	})

	//handler是处理服务器，写程127.0.0.1的话只支持本机，写成0.0.0.0就可以支持任何ip
	err := http.ListenAndServe("0.0.0.0:7222", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("服务启动成功在127.0.0.1:7222")
}
