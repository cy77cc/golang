package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var rwm sync.RWMutex
var do = make(chan bool, 5)

func readDB() {
	do <- true
	rwm.RLock()
	<-time.After(2 * time.Second)
	fmt.Println("读数据库")
	rwm.RUnlock()
	wg.Done()
	<-do
}

func writeDB() {
	do <- true
	rwm.Lock()
	ticker := time.NewTicker(time.Second)
	for i := 0; i < 2; i++ {
		<-ticker.C
	}
	fmt.Println("写数据库")
	rwm.Unlock()
	wg.Done()
	<-do
}

func main() {
	wg.Add(40)
	for i := 0; i < 30; i++ {
		go readDB()
	}

	for i := 0; i < 10; i++ {
		go writeDB()
	}
	wg.Wait()
}
