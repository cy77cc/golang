package main

import (
	"fmt"
	"time"
)

func main() {
	myChan := make(chan int, 10)
	go func() {
		for i := 0; i <= 10; i++ {
			myChan <- i
		}
	}()
	// 管道遍历没有下标
	go func() {
		// 如果管道没有关闭，会永远尝试进行下一次读取
		for v := range myChan {
			fmt.Println(v)
		}
	}()
	time.Sleep(2*time.Second)
}
