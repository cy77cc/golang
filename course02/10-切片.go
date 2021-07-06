package main

import (
	"fmt"
	"sort"
)

func main() {
	//var slice = []int{10, 20, 30}
	//fmt.Println(slice)

	//var slice = make([]int, 10, 20)
	//fmt.Println(slice)
	//slice = append(slice, 10)
	//fmt.Println(slice)
	//fmt.Println(cap(slice))
	//for i := 0; i < 8; i++ {
	//	slice[i] = i
	//}
	//fmt.Println(slice)

	var slice = []int{12, 23, 47}
	var slice2 = []int{24, 94, 84}

	// ... 就相当于把切片里面的值都取出来
	slice = append(slice, slice2...)
	var slice3 []int = make([]int, 10, 20)
	copy(slice3, slice)
	fmt.Println(slice)
	fmt.Println(slice3)
	sort.Ints(slice)
	fmt.Println(slice)
}
