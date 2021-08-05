package main

import (
	"fmt"
	"sync"
	"time"
)

//读写锁API介绍
//多路只读
//一路只写
//读写互斥
func main02() {
	var rwm sync.RWMutex

	//锁定为写模式------ 一路只写不可读
	rwm.Lock()
	//释放读写锁
	rwm.Unlock()

	//锁定为读模式----- 多路只读
	rwm.RLock()

	//释放读写锁
	rwm.RUnlock()
}

/*
数据库的一写多读
ReadDB方法设定为多路只读
WriteDB方法设定为单路只写
创建5读5写10条协程,观察读写锁效果
*/
func main() {
	var wg sync.WaitGroup

	//声明读写锁
	var rwm sync.RWMutex

	for i := 0; i < 5; i++ {
		wg.Add(2)
		go func() {

			//锁定为只读模式,允许多个协程同时抢到读锁
			rwm.RLock()
			fmt.Println("读取数据库")
			<-time.After(2 * time.Second)

			//解锁
			rwm.RUnlock()
			wg.Done()
		}()

		go func() {
			//锁定为写模式,只允许一条协程抢到锁
			rwm.Lock()
			fmt.Println("写入数据库")
			<-time.After(2 * time.Second)

			//解锁
			rwm.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
}
