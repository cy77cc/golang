package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

var (
	regImgStr = `<img[^>]+?src="(.+?)"`
	//<a class="page-next" href="/tupian/index_2.html">下一页</a>
	regFenye = `<a class=.page-next. href=.(.+?)('|")`
	//	信号量管道
	chSem = make(chan int, 5)
)

var count int

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

func imgSpider(src string) {
	var wg sync.WaitGroup
	html := getPageHtml(src)
	//fmt.Println(html)
	regImg := regexp.MustCompile(regImgStr)
	res := regImg.FindAllStringSubmatch(html, -1)
	//fmt.Println(res)
	for _, v := range res {
		wg.Add(1)
		go func(url string) {
			chSem <- 1
			resp, _ := http.Get(url)
			body := resp.Body
			defer func() {
				_ = body.Close()
			}()
			bytes, _ := ioutil.ReadAll(body)
			split := strings.Split(url, "/")
			_ = ioutil.WriteFile(`E:\golang\正则表达式\img\`+split[len(split)-1], bytes, 0666)
			fmt.Println(split[len(split)-1] + "爬取成功")
			<-chSem
			wg.Done()
		}("https:" + v[1])
	}
	wg.Wait()
}

func getPage(src string) {
	count++
	html := getPageHtml(src)
	regFen := regexp.MustCompile(regFenye)
	res := regFen.FindAllStringSubmatch(html, -1)
	imgSpider(`https://www.ivsky.com/` + res[0][1])
	if count == 100 {
		fmt.Println("爬取完毕……")
		return
	}
	getPage(`https://www.ivsky.com/` + res[0][1])
}

func main() {
	getPage(`https://www.ivsky.com/tupian/`)
}
