package main

import (
	"fmt"
	"time"
)

func add(n int) int {
	if n == 1 {
		return 1
	}
	return n + add(n-1)
}

func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fbn(n-1) + fbn(n-2)
}

func main() {
	starttime := time.Now().UnixNano()
	var sum = add(10)
	fmt.Println(sum)
	var num = fbn(30)
	fmt.Println(num)
	endtime := time.Now().UnixNano()
	fmt.Println(endtime - starttime)
}
