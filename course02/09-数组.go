package main

import "fmt"

func main() {
	//var arr [3]int
	//arr[0] = 10
	//arr[1] = 20
	//arr[2] = 40
	//for _, v := range arr {
	//	fmt.Println(v)
	//}

	var arr [3][4]int
	arr[0] = [4]int{1, 2, 3, 4}
	arr[1] = [4]int{5, 6, 7, 8}
	arr[2] = [4]int{9, 10, 11, 12}
	for i, v := range arr {
		fmt.Println(i, v)
	}

	fmt.Println(arr[1][2])
}
