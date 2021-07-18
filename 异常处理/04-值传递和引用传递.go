package main

import (
	"encoding/json"
	"fmt"
)

//type person struct {
//	age int
//}

func base(a int) {
	a = 20
	//fmt.Println("函数改变后", a)
}

func array(arr [4]int) {
	//fmt.Println(arr)
	arr[3] = 50
	//fmt.Println(arr)
}

func slice(sli []int) {
	sli[2] = 200
	//sli = append(sli, 1, 2, 3, 4, 5, 6)
	//fmt.Println(sli)
}

func maps(m map[string]string) {
	m["姓名"] = "李四"
}

func main() {
	var a int = 10
	fmt.Println("基本数据类型调用函数之前", a)
	base(a)
	fmt.Println("基本数据类型调用函数之后", a)

	var arr = [...]int{10, 20, 39, 40}
	fmt.Println("数组调用函数之前", arr)
	array(arr)
	fmt.Println("数组调用函数之后", arr)

	// 切片底层是一个结构体，当参数传递时，把结构体传入，引用数组的地址是没有变化的，所以在函数内容改变slice的内容会影响原来的内容
	ints := make([]int, 0)
	ints = append(ints, 10, 20, 30)
	fmt.Println("切片调用函数之前", ints)
	slice(ints)
	fmt.Println("切片调用函数之后", ints)

	var m map[string]string
	m = make(map[string]string)
	m["姓名"] = "张三"
	m["年龄"] = "20"
	fmt.Println("map函数调用前", m)
	maps(m)
	fmt.Println("map函数调用后", m)
	marshal, _ := json.Marshal(m)
	fmt.Println(string(marshal))
}
