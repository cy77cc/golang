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
	regPhoneStr = `(>)(1[3864579][0-9]{9})(<)`
)

func main() {
	html := getPageHtml(`http://www.taohaoba.com/`)
	//fmt.Println(html)
	regPhone := regexp.MustCompile(regPhoneStr)

	res := regPhone.FindAllStringSubmatch(html, -1)
	//fmt.Println(res)
	for _, v := range res {
		fmt.Println(v[2])
	}
}
