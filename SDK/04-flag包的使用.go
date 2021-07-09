package main

import "flag"

func main() {
	//args := os.Args
	//fmt.Println(args)
	//for _, v := range args {
	//	fmt.Println(v)
	//}

	//namePtr := flag.String("name", "无名氏", "用户名")
	//agePtr := flag.Int("age", 0, "年龄")
	//moneyPtr := flag.Float64("money", 0, "存款")
	//isStupid := flag.Bool("isstupid", true, "是否愚蠢")
	//
	// 解析参数
	//flag.Parse()
	//fmt.Println(*namePtr, *agePtr, *moneyPtr, *isStupid)

	var name string
	flag.StringVar(&name, "name", "无名氏", "用户名")
}
