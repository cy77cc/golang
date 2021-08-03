package main

import (
	"fmt"
	"time"
)

// 校验管道数据的有效性
func main01() {
	/*
	ch称谓一个未满且关闭的管道
	*/
	ch := make(chan int, 3)
	ch <- 123
	ch <- 456
	close(ch)

	//已经写入的数据能够有效被读出
	x, ok := <-ch
	fmt.Println(x, ok)

	x, ok = <-ch
	fmt.Println(x, ok)

	//未被写入的数据时零值数据，ok为false
	x, ok = <-ch
	fmt.Println(x, ok)
}

// 关闭管道
func main() {
	ch := make(chan int, 5)

	go func() {
		for v := range ch {
			fmt.Println(v)
		}
		fmt.Println("goroutine over")
	}()

	var n int
	for {
		ch <- 123
		n++
		if n > 10 {
			close(ch)
			break
		}
	}
	time.Sleep(1*time.Second)
}