package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 1; i <= 10; i++ {
		go func(i int) {
			for j := 1; j <= 10; j++ {
				if i == 5 {
					/*
						使当前go程放弃处理器，以让其它go程运行。它不会挂起当前go程，因此当前go程未来会恢复执行。
					*/
					runtime.Gosched()
				}
				fmt.Printf("协程%d:%d\n", i, j)
			}
		}(i)
	}
	time.Sleep(2 * time.Second)
}
