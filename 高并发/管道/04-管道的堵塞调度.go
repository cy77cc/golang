package main

import (
	"fmt"
	"time"
)

// 创建一个三缓存的管道
var intChan = make(chan int, 3)

func count(grName string, n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(grName, i)
		time.Sleep(200 * time.Millisecond)
	}
	intChan <- 0
}

func main() {
	/*
		三个协程分别数数
		要求主协程恰好在所有协程结束时结束
	*/
	go count("didi", 10)
	go count("meimei", 15)
	go count("diedie", 17)

	for {
		if len(intChan) == 3 {
			break
		}
	}

}
