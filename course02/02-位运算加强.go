package main

import "fmt"

func main() {
	var b int8 = 23
	// 有符号的数移位之后是补码
	fmt.Println(b<<3)
}
