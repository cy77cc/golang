package main

/*
从今日头条.json中，解析获取每条新闻，提取其title，label，content，封装到News对象中
最总打印新闻对象切片
再将新闻兑现切片序列化为json
再讲新闻对象切片编码到《新闻.json》
再讲上述json反序列化为map
将《新闻.json》解码为切片嵌套结构体
*/

type New struct {
	Title string `json:"title"`
	Label []string `json:"label"`
}

func main() {

}
