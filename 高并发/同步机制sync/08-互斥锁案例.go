package main

import (
	"fmt"
	"sync"
	"time"
)

/*
银行账户案例
	创建银行账户类
	存取款方法需要并发安全
	查询余额和打印流水不需要
	创建账户实例，并发执行存取款，查询余额，打印流水操作
*/

type Account struct {
	money float64
	//账户的互斥锁
	mt sync.Mutex
}

//存钱
func (a *Account) SaveMoney(n float64) {
	//想存钱先抢锁
	a.mt.Lock()
	fmt.Println("存钱中……")
	<-time.After(2 * time.Second)
	a.money += n
	fmt.Println("存钱成功", n)
	a.Query()
	//存钱之后释放同步锁
	a.mt.Unlock()
}

//取钱
func (a *Account) GetMoney(n float64) {
	a.mt.Lock()
	fmt.Println("取钱……")
	<-time.After(2 * time.Second)
	a.money -= n
	fmt.Println("取钱成功", n)
	a.Query()
	a.mt.Unlock()
}

//查询余额
func (a *Account) Query() {
	<-time.After(2 * time.Second)
	fmt.Println("当前余额：", a.money)
}

//打印流水
func (a *Account) PrintHistory() {
	<-time.After(2 * time.Second)
	fmt.Println("打印流水")
}

func main() {
	var wg sync.WaitGroup
	account := Account{1000, sync.Mutex{}}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			account.SaveMoney(100)
			wg.Done()
		}()
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			account.GetMoney(50)
			wg.Done()
		}()
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			account.Query()
			wg.Done()
		}()
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			account.PrintHistory()
			wg.Done()
		}()
	}

	wg.Wait()
}
