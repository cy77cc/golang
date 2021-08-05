package main

import (
	"fmt"
	"time"
)

/*
循环从一写两读三条管道中随意选择一条能走的路
等所有路都走不通了就退出循环
*/

func main() {
	chA := make(chan int, 5)
	chB := make(chan int, 4)
	chB <- 123
	chB <- 123
	chB <- 123
	chB <- 123

	chC := make(chan int, 3)
	chC <- 124
	chC <- 124
	chC <- 124

	//go taskA(chA)
	//go taskB(chB)
	//go taskC(chC)

OUTER:
	for {
		select {
		case chA <- 123:
			fmt.Println("执行任务A")
			time.Sleep(time.Second)
		case <-chB:
			fmt.Println("执行任务B")
			time.Sleep(time.Second)
		case <-chC:
			fmt.Println("执行任务C")
			time.Sleep(time.Second)
		default:
			fmt.Println("全部任务已结束")
			break OUTER
		}
	}
}

func taskA(ch chan int) {
	for {
		fmt.Println("taskA")
		ch <- 123
		time.Sleep(time.Second)
	}
}

func taskB(ch chan int) {
	for {
		fmt.Println("taskB")
		<-ch
		time.Sleep(time.Second)
	}
}

func taskC(ch chan int) {
	for {
		fmt.Println("taskC")
		<-ch
		time.Sleep(time.Second)
	}
}
