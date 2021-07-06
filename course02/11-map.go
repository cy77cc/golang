package main

import "fmt"

func main() {
	var scoreMap = map[string]int{
		"nake":  20,
		"alice": 10,
		"jack":  30,
	}
	fmt.Println(scoreMap)
	fmt.Println(len(scoreMap))
	fmt.Println(scoreMap["nake"])
	for key, value := range scoreMap {
		fmt.Println(key, value)
	}
	// 这样声明的map不能赋值，要make或者赋一个空值，因为map是引用类型，不赋值或make没有地址
	//var tempMap map[string]float64
}
