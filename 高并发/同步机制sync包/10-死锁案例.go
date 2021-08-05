package main

import (
	"sync"
)

var wg sync.WaitGroup

/*AB相互要求对方先发红包自己再发*/
func main11() {
	chA := make(chan int)
	chB := make(chan int)

	// a
	go func() {
		<-chA
		chB <- 123
	}()

	// b
	<-chB
	chA <- 123
}

func main12() {
	chA := make(chan int)
	chB := make(chan int)

	// a
	go func() {
		wg.Add(1)
		<-chA
		chB <- 123
		wg.Done()
	}()

	// b
	go func() {
		wg.Add(1)
		<-chB
		chA <- 123
		wg.Done()
	}()
	wg.Wait()
}

/*读写协程相互锁定对方需要的资源*/
func main() {

	var rwm sync.RWMutex

	ch := make(chan int)

	wg.Add(2)
	go func() {
		//锁定为只读,排斥写入协程
		rwm.RLock()
		//需要写入协程,没协程写就永远读不出来
		<- ch
		rwm.RUnlock()
		wg.Done()
	}()

	go func() {
		//锁定为只写,排斥读取协程
		rwm.Lock()
		//需要读取协程,没协程读就永远写不进去
		ch <- 123
		rwm.Unlock()
		wg.Done()
	}()
	wg.Wait()
}
