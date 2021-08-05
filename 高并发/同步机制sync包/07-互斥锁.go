package main

import (
	"fmt"
	"sync"
)

// 使用10个协程频繁修改一个数据，演示何为“并发不安全”
func main01() {
	//声明同步锁
	var wg sync.WaitGroup
	var money = 2000
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				money += 1
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(money)
}

//演示加同步锁后的安全访问场景
func main() {

	//不允许并发访问，只允许排队同步访问
	var wg sync.WaitGroup
	var mtx sync.Mutex

	/*
		锁住（一次只能被一个协程锁住，其余想要抢到锁的协程阻塞等待至前面的协程将锁释放
		mtx.Lock()的可能性有两种：
			1. 抢到锁，继续向下执行
			2. 没抢到，阻塞等待至前面的协程将锁释放
	*/
	//mtx.Lock()

	//解锁，解锁释放之后其他抢这把锁的协程就会得到抢锁机会
	//mtx.Unlock()

	var money = 2000
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Printf("协程%d开始抢锁\n", n)
			//抢锁，如果成功就向下执行，
			//否则阻塞等待至抢到为止（抢到的协程不释放，就一直阻塞)，
			//所有抢锁的协程都是资源竞争者
			mtx.Lock()
			fmt.Printf("协程%d抢到锁\n", n)
			for j := 0; j < 10000; j++ {
				money += 1
			}
			//数据操作完毕将锁释放
			mtx.Unlock()
			fmt.Printf("协程%d将锁释放\n", n)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(money)
}
