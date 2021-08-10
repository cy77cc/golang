package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	//reEmailStr = `[1-9]\d{4,9}@qq\.com`
	//reEmailStr = `([1-9]\d{4,})@(qq\.com)` // 加括号分组
	reEmailStr = `([a-z\d_]+)@([a-z0-9]+?\.[a-z]+(\.[a-z]+)?)` // 加括号分组
)

func spiderEmail() {
	//获得页面html（庞大的字符串）
	html := getPageHtml(`https://tieba.baidu.com/p/2366935784`)

	//编译正则字符串，得到正则表达式对象
	reEmail := regexp.MustCompile(reEmailStr)

	//从html中寻找所有匹配的子串，及子串内的小分组
	//-1代表查找所有
	res := reEmail.FindAllStringSubmatch(html, -1)
	fmt.Println(res)
	for _, v := range res {
		fmt.Println("邮箱", v[0])
		fmt.Println("号码", v[1])
		fmt.Println("主机名", v[2])
		fmt.Println()
	}
}

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
