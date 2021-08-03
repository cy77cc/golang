package main

import (
	"fmt"
	"strconv"
	"time"
)

type product1 struct {
	name string
}

var chQuit = make(chan interface{})

func main() {

	chStorage := make(chan product1, 5)

	chShop := make(chan product1)

	go producer(chStorage)

	go logistics(chStorage, chShop)

	go consumer(chShop)

	msg := <-chQuit
	fmt.Println(msg)
}

// 生产者
func producer(chStorage chan<- product1) {
	for {
		p := product1{"商品" + strconv.Itoa(time.Now().Second())}
		time.Sleep(time.Second)

		//把产品丢进库存
		chStorage <- p

		fmt.Println("生产了", p.name)
	}
}

// 物流公司
func logistics(chStorage <-chan product1, chShop chan<- product1) {
	for {
		p := <-chStorage
		chShop <- p
		fmt.Println("转运了", p.name)
	}
}

// 消费者
func consumer(chShop <-chan product1) {
	for {
		p := <-chShop
		fmt.Println("消费了", p.name)

		//if count >= 10 {
		//	break
		//}
		quitTime := time.Now().Format("15:04:05")
		if quitTime == "10:18:00" {
			fmt.Println(quitTime)
			chQuit <- "爷罢工了"
		}
	}

}
