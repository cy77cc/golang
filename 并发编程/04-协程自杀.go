package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		for i := 1; i <= 10; i++ {
			if i == 5 {
				fmt.Println("goroutine: 还我头来！！！")
				// 当 i == 5 时杀死协程
				runtime.Goexit()
			}
			fmt.Println("go", i)
			time.Sleep(500 * time.Millisecond)
		}
	}()
	for i := 1; i <= 10; i++ {
		fmt.Println("main", i)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("main over")
}
