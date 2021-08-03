package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// 向等待组中添加一个协程（注册）
	//wg.Add()

	// 向等待组中减掉一个协程（注销）
	//wg.Done()

	// 阻塞等待组中的协程数归零
	//wg.Wait()

	wg.Add(1)
	go func() {
		fmt.Println("协程A开始工作")
		time.Sleep(3 * time.Second)
		fmt.Println("协程A over")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("协程B开始工作")
		timer := time.NewTimer(5 * time.Second)
		<-timer.C
		fmt.Println("协程B over")
		timer.Stop()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("协程C开始工作")

		ticker := time.NewTicker(1 * time.Second)
		for i := 0; i < 4; i++ {
			<-ticker.C
		}
		ticker.Stop()
		fmt.Println("协程C over")
		wg.Done()
	}()

	wg.Wait()
}
