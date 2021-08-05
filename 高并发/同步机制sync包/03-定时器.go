package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)
	start := time.Now().UnixNano()
	fmt.Println("计时开始")
	<-timer.C
	fmt.Println("计时结束", time.Now().UnixNano()-start)
	fmt.Println("bombombom")

	fmt.Println("================================")
	fmt.Println("计时开始")
	<-time.After(3*time.Second)
	fmt.Println("bombombom")
	fmt.Println("计时结束")
}
