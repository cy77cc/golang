package main

import (
	"fmt"
	"math"
	"time"
)

/*
100条协程并发求1-10000平方根
最大并发数控制在5
管道实现
*/

func getSqrt(name string, n int, semaphore chan string) {
	// 想执行，先注册
	// 能写入就执行，写不进去就阻塞到能写入为止
	semaphore <- name

	res := math.Sqrt(float64(n))
	fmt.Printf("%d的平方根是%.2f\n", n, res)

	time.Sleep(time.Second)

	// 任务执行完毕，从信号量控制管道注销自己，以便为其他并发任务腾出空间
	<-semaphore
}

func main() {
	// 信号量控制管道
	// 凡要并发执行的协程必须先讲协程名字注册到该管道才能执行
	semaphore := make(chan string, 5)
	for i := 1; i <= 100; i++ {
		getSqrt(fmt.Sprintf("名字%d", i), i, semaphore)
	}

	for {
		time.Sleep(time.Second)
	}

}
