package main

import "fmt"

func main() {
	var myChan chan int
	/*
		只读管道
		var myChan <-chan int
		只写管道
		var myChan chan<- int
	*/
	myChan = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		myChan <- i
	}
	for {
		if len(myChan) == 0 {
			break
		}
		fmt.Println("长度", len(myChan))
		x, ok := <-myChan
		if ok {
			fmt.Println("读到正确的值", x)
		} else {
			fmt.Println("无法正确进行读取")
		}
	}
	/*
		close关闭一个管道
		关闭的管道，无法再写入数据，但可以读出
		一个已经关闭且读空的管道，继续读读出零值
		关闭一个未初始化的管道会panic
		关闭管道会产生一个广播机制，所有向管道读取消息的goroutine都会收到消息
	*/
	close(myChan)
}
