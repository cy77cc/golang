package main

import (
	"fmt"
	"sync"
	"time"
)

func condAPI() {
	//要监听的变量
	bitcoinRising := false

	//创建条件变量
	cond := sync.NewCond(&sync.Mutex{})

	//市场协程
	go func() {
		for {

			//枷锁修改为涨，并通知投资者
			cond.L.Lock()
			bitcoinRising = true
			cond.Broadcast()
			cond.L.Unlock()

			//持续涨
			time.Sleep(time.Second)

			//枷锁修改为跌
			cond.L.Lock()
			bitcoinRising = false
			cond.L.Unlock()

			//持续一段时间
			time.Sleep(2 * time.Second)
		}
	}()

	//投资者协程
	for {
		//只要比特币跌，就停止投资，等待涨的消息（释放锁）
		//加锁监听比特币的变化
		cond.L.Lock()
		for !bitcoinRising {
			fmt.Println("停止比特币投资")
			//比特币没开始涨，就释放锁，等待比特币开始涨的消息

			//阻塞等待消息
			//收到涨的消息，就继续向下执行（其他协程通过cond.Signal/Broadcast()发送涨跌消息）
			cond.Wait() // 内部把锁释放

			//收到了市场变化的消息后继续执行
		}

		//走到这里说明bitcoinRising==true
		fmt.Println("开始投资比特币...")
		cond.L.Unlock()
	}

}

func testCond() {
	cond := sync.NewCond(&sync.Mutex{})
	condition := false

	go func() {
		time.Sleep(time.Second)
		cond.L.Lock()

		fmt.Println("[1] 变更condition状态，并发出变更通知")
		condition = true
		cond.Signal() // 通知一条协程
		//cond.Broadcast() 广播，通知所有协程
		fmt.Println("[1] 继续后续处理")

		cond.L.Unlock()
	}()

	cond.L.Lock()

	fmt.Println("[2] condition.............1")
	for !condition {
		fmt.Println("[2] condition.............2")
		//等待Cond消息通知（释放锁？）
		cond.Wait()
		fmt.Println("[2] condition.............3")
	}
	fmt.Println("[2] condition.............4")
	cond.L.Unlock()
}

func main() {
	testCond()
	condAPI()
}
