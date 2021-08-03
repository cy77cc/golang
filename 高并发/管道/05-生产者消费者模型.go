package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
生产者每天生产一件产品
生产者每生产一件，消费者就消费一件
消费10轮后，主协程退出
*/

/*
思路：
	生产者源源不断地向【商店管道】写入产品
	消费者阻塞等待从【商店管道】读出数据
	消费者每读出一次数据（消费一次）就向【计数管道】写入一个数据
	主协程阻塞从【计数管道】读出10个数据就结束
*/

type product struct {
	name string
}

func main() {
	// 创建【商店管道】【计数管道】
	chanShop := make(chan product, 5)
	chanCount := make(chan int, 5)

	go produce(chanShop)
	go consume(chanShop, chanCount)

	for i := 0; i < 10; i++ {
		<- chanCount
	}

}

//消费者阻塞等待从【商店管道】读出数据
//消费者每读出一次数据（消费一次）就向【计数管道】写入一个数据
func consume(chanShop <-chan product, chanCount chan<- int) {
	for {
		p := <-chanShop
		fmt.Println("消费者吃掉了", p.name)

		//	报个数
		chanCount <- 1
	}
}

//生产者源源不断地向【商店管道】写入产品
func produce(chanShop chan<- product) {
	for {
		p := product{"产品" + strconv.Itoa(time.Now().Second())}
		fmt.Println("生产者向商店输送了", p.name)
		chanShop <- p
		time.Sleep(1 * time.Second)
	}
}
