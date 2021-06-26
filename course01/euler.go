package main

import (
	"fmt"
	"math"
)

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

func main() {
	//c := 3 + 4i
	//b := cmplx.Abs(c)
	//fmt.Println(b)
	//fmt.Println(math.E)
	//triangle()
	const (
		// 会延续表达式到下一个变量
		b = 1 << (10 * iota)
		kb // 1 << (10 * iota) iota = 1
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
