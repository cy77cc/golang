package main

import (
	"fmt"
	"time"
)

/*
cooperate	合作
routine		日程，事物
coroutine	协程（可以相互协作的并发任务）
goroutine	go语言协程，go程
*/

func newTask(i int) {
	for {
		fmt.Println("我是子协程", i)
		time.Sleep(1*time.Second)
	}
}

func main() {
	for i := 1; i <= 10000; i++ {
		go newTask(i)
	}
	for {
		fmt.Println("我是主程序")
		time.Sleep(1*time.Second)
	}
}
