package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("计时开始", time.Now())
	go func() {
		time.Sleep(3*time.Second)
		timer.Stop()
		fmt.Println("炸弹已拆除")
		//以当前时间为基准（而不是timer被创建的时间），将定时器重置
		timer.Reset(0)
	}()

	//阻塞5秒
	//select {
	//case t := <-timer.C:
	//	fmt.Println("爆炸了，世界末日了", t)
	//default:
	//	fmt.Println("炸弹拆除了")
	//}
	t := <- timer.C
	fmt.Println("炸弹已拆除", t)
}
