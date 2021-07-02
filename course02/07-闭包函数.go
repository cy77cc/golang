package main

import "fmt"

func study(name string) func(hours int) {
	var progress int
	return func(hours int) {
		progress += hours
		fmt.Printf("%s学习了%d小时\n", name, hours)
		fmt.Printf("%s学习进度是%d\n", name, progress)
	}
}

func main() {
	likuiStudy := study("李逵")
	likuiStudy(10)
	likuiStudy(20)
	wusonStudy := study("武松")
	wusonStudy(30)
	wusonStudy(40)
	wusonStudy(90)
}
