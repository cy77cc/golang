package main

import "fmt"

// 无参函数
func add0() {
	fmt.Println("add0")
}

// 有参函数
func add1(arg int) {
	fmt.Println("add1", arg)
}

// 多个参数，同种类型
func add2(a, b int) {
	fmt.Println(a + b)
}

// 不定长产生参数
//args 接受任意多个整型参数是一个整型切片
func add5(args ...int) {
	fmt.Printf("args的类型是%T\n", args)
	fmt.Println(args)
	fmt.Println("args的长度是", len(args), cap(args))
	for i, v := range args {
		fmt.Printf("i = %d, v = %d\n", i, v)
	}
}

func calculate(a, b int) (sum, sub, mul, div int) {
	sum = a + b
	sub = a - b
	mul = a * b
	div = a / b
	return
}

var (
	likuiProgress int
	wusonProgress int
)

// 匿名函数
func getStudy(name string) func(int) {
	study := func(hours int) {
		fmt.Printf("%s学习了%d小时\n", name, hours)
		if name == "李逵" {
			likuiProgress += hours
		} else if name == "武松" {
			wusonProgress += hours
		}
	}
	return study
}

func main() {
	//add5(1, 3, 3, 6)
	for {
		go func() {
			fmt.Println("go")
		}()
	}
	sum, sub, mul, div := calculate(10, 5)
	fmt.Println(sum, sub, mul, div)
	func() {
		fmt.Println("yes")
	}()
}
