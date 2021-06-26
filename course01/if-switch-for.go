package main

import (
	"fmt"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func main() {
	if a := 10; a == 10 {
		fmt.Println("yes")
	}
	//var score int
	//fmt.Scanf("%d", &score)
	//switch {
	//case score >= 80:
	//	fmt.Println("良好")
	//case score >= 60:
	//	fmt.Println("优秀")
	//case score < 60:
	//	fmt.Println("不及格")
	//}

	fmt.Println(convertToBin(30))
}
