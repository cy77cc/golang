package main

import (
	"fmt"
	"math/cmplx"
)

func variable() {
	var a int
	var s string
	fmt.Println(a, s)

	var b, c int = 3, 4
	var s2 string = "abc"
	fmt.Println(b, c, s2)
	/*
		bool,string
		int, int8, int16, int32, int64, uintptr
		byte, rune
		float32, float64, complex64, complex128
	*/
}

func euler() {
	// c := 3 + 4i
	// fmt.Println(cmplx.Abs(c))
	fmt.Println(cmplx.Exp(10))
}

func enums() {
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, pb)
}

func main() {
	// variable()
	// euler()
	enums()
	fmt.Println()
}
