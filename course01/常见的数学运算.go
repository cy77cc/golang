package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		加减乘除，求余
	*/
	fmt.Println("5 + 3 =", 5+3)
	fmt.Println("5 - 3 =", 5-3)
	fmt.Println("5 * 3 =", 5*3)
	fmt.Println("5 / 3 =", 5/3)
	fmt.Println("5 % 3 =", 5%3)

	// 乘方和开方
	fmt.Println("5 ^ 3 =", math.Pow(5, 3))
	fmt.Println("125开三次方 =", int(math.Pow(125.0, 1.0/3)))

	// 四舍五入
	/*
		负数的四舍五入原则，先对绝对值四舍五入，再加负号
	*/
	fmt.Println("3.49的四舍五入 =", math.Round(3.49))
	fmt.Println("3.51的四舍五入 =", math.Round(3.51))
	// 向下取整
	fmt.Println("3.99的四舍五入 =", math.Floor(3.99))
	// 向上取整
	fmt.Println("3.01的四舍五入 =", math.Ceil(3.01))

	// 绝对值
	fmt.Println("-3.14的绝对值", math.Abs(-3.14))

	// 三角函数,参数必须是弧度而不是角度
	fmt.Println("30°的正弦 =", math.Sin((30.0/180)*math.Pi))
	fmt.Println("30°的余弦 =", math.Cos((30.0/180)*math.Pi))
	fmt.Println("30°的正切 =", math.Tan((30.0/180)*math.Pi))
	fmt.Println("30°的余切 =", 1.0/math.Tan((30.0/180)*math.Pi))

	// 反三角函数
	fmt.Println("正弦为0.5的角度是", math.Asin(0.5)*360)
}
