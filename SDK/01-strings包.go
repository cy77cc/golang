package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "html"

	// 字符串中是否包含子串
	fmt.Println(strings.Contains(str, "h"))

	// 字符串中是否包含字符
	fmt.Println(strings.ContainsRune(str, 'h'))

	// 字符创中是否包含参数字符串中任一字符
	fmt.Println(strings.ContainsAny(str, "hm"))
	//fmt.Println("Hello, World")
	//fmt.Printf("%c", 123)

	// 子串在字符串中第一次出现的位置，没有的话返回-1
	fmt.Println(strings.Index(str, "tm"))

	//	一个中文占三个字节，一个汉字所占的数组容量也是三个
	// 你妹啊  你 0-2   妹 3-5  啊 6-8
	fmt.Println(strings.Index("你妹啊", "妹")) // 3

	// 返回chars里面任一字符在s中第一次出现的位置
	fmt.Println(strings.IndexAny("html", "bt"))

	//	转换大小写
	fmt.Println(strings.ToLower("HeLLo"))
	fmt.Println(strings.ToUpper("HeLLo"))
	fmt.Println(strings.Title("Hello world"))

	// 比较首字母在字符集中出现的位置 ASCII大小
	// a == b 0
	// a > b 1
	// a < b -1
	fmt.Println(strings.Compare("hello", "world")) // -1

	//var str1 string
	//for i := 0; i < 100; i++ {
	//	str1 += "-"
	//	fmt.Print("\r")
	//	fmt.Print(str1 + ">")
	//	time.Sleep(1 * time.Second)
	//}

	// 裁剪
	// 第二个参数f 自定义裁剪规则（对传入的参数字符返回true代表要裁剪掉）
	fmt.Println(strings.TrimFunc("男：你妹啊 女：你大爷的 合：我们是吉祥的一家男", func(r rune) bool {
		if r == '男' || r == '女' {
			return true
		}
		return false
	}))

	// 分割
	fmt.Println(strings.Split("we are student", " "))
	fmt.Println(strings.SplitAfter("we,are,student", ","))

	// 拼接
	fmt.Println(strings.Join([]string{"男", "女", "娘炮"}, "*"))
}
