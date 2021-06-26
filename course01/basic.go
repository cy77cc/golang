package main

import "fmt"

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

func main() {
	variable()
	fmt.Println()
}
