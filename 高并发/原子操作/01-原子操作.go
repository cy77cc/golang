package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
用原子操作来替换锁，其主要原因是：
原子操作由底层硬件支持，而锁则由操作系统提供API实现
若实现相同的功能，原子操作通常会更有效率

只有操作基本数据类型
*/


//什么是原子操作
func main01() {
	var a int64 = 23

	var mt sync.Mutex


	val := atomic.LoadInt64(&a)
	fmt.Println(val)

	mt.Lock()
	fmt.Println(a)
	mt.Unlock()
}

func main() {
	var a int64 = 20

	//保证将456赋值到a的地址中（期间a一定不会被其他人访问
	//atomic.StoreInt64(&a, 456)
	//fmt.Println(a)


	//newVal := atomic.AddInt64(&a, 30)
	//fmt.Println(newVal)
	//fmt.Println(a)

	//oldVal := atomic.SwapInt64(&a, 65)
	//fmt.Println(oldVal, a)

	//假如第二个参数old和真实的值不一致就会交换失败
	//确保a在没有并发修改的情况下被重新赋值为456
	swapped := atomic.CompareAndSwapInt64(&a, 123, 456)
	fmt.Println(swapped, a)
}
