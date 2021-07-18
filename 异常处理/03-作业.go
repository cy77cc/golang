package main

import "fmt"

/*
封装给一个human类（姓名、年龄、性别、身高、资产、颜值）
为human类封装一个结婚方法；func (h *human)marry(other *human) error
为human类封装一个设置征婚标准方法，用于设置能接受的配偶条件：func (h *human)setStandard(minHeight, maxHeight, minAge, maxAge, minLooking, minMoney) error
如果性别与自己想同，panic
如果征婚对象其他方面达不到自己设置的要求，就返回自定义的marryError，将信息尽可能完整的提出来
主函数，创建一个女神，创建一堆男士，令他们尝试结婚，并打印试婚结果
*/

type human struct {
	name    string
	age     int
	male    bool
	height  float64
	money   float64
	looking int
}

type marryError struct {
	age int
	height float64
	looking float64
	money float64
}

func (m *marryError) Error() string {
	return ""
}

func (h *human) marry(other *human) error {
	return nil
}

func (h *human) setStandard(minAge, maxAge int, minHeight, maxHeight, minLooking, minMoney float64) error {
	return nil
}

func main() {
	var arr = [...]int{10, 20}
	// 一个字节加一个地址
	fmt.Printf("第一个元素的地址%p, 第二个元素的地址%p", &arr[0], &arr[1])
}
