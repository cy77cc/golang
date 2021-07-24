package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	//response, err := http.Get(`http://www.baidu.com`)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	////响应体是IO资源，使用完毕要释放
	//defer func() {
	//	_ = response.Body.Close()
	//}()
	////读取响应体的内容
	//bytes, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(bytes))

	//创建一个io.Reader，作为请求体
	body := strings.NewReader("rmb=0.5&hobby=杰克和韩寒")
	resp, err := http.Post("http://httpbin.org/post", "application/x-www-form-urlencoded", body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
}
