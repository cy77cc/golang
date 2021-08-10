package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func getPageHtml(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("err", err)
	}
	body := resp.Body
	defer func() {
		_ = body.Close()
	}()
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err)
	}
	html := string(bytes)
	return html
}

var (
	regUrlStr = `href="((https?://.+?)|(.+?))"` // 使用非贪婪模式匹配到“就结束
)

func main() {
	html := getPageHtml(`https://www.jxust.edu.cn/`)
	//fmt.Println(html)
	regUrl := regexp.MustCompile(regUrlStr)
	res := regUrl.FindAllStringSubmatch(html, -1)
	//fmt.Println(res)
	for _, v := range res {
		fmt.Println(v[1])
	}
}
