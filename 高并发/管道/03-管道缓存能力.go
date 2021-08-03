package main

import "fmt"

func main() {
	myChan := make(chan int, 5)
	myChan <- 123

	fmt.Println("管道的长度是", len(myChan))
	fmt.Println("管道的容量是", cap(myChan))
	//	已满的管道无法再进行写入
	myChan <- 124
	myChan <- 125
	myChan <- 126
	myChan <- 127
	//myChan <- 128 报错
	//  已读空的管道无法读出
	<- myChan
	<- myChan
	<- myChan
	<- myChan
	<- myChan
	//<- myChan 已经读空了会报错

	fmt.Println("main over!!!")
}
