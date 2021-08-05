package main

import (
	"fmt"
	"sync"
	"time"
)

/*
服务器负载控制
监听最大客户端数量
服务端协程：只要服务端客户端数量达到阙值，就通知控制协程，就进入阻塞等待
监察协程：收到服务端预警，削减客户端数量，并通知服务端（过载预警已解除）
*/

const MaxConnections = 5

type server struct {
	connections int
	chAlarm     chan bool
	cond        *sync.Cond
}

func (s *server) serve() {
	go func() {
		for {
			//阻塞监听是否有预警
			<-s.chAlarm
			fmt.Println("客户端数量过量")
			<-time.After(time.Second)

			//枷锁削减客户端数量，并发送预警解除广播
			s.cond.L.Lock()
			s.connections--

			//发送预警解除广播
			s.cond.Signal()
			s.cond.L.Unlock()
		}
	}()

	for {
		s.cond.L.Lock()

		//监听是否过载
		for s.connections == MaxConnections {
			// 发送预警
			s.chAlarm <- true
			//阻塞等待预警解除
			s.cond.Wait()
		}
		//接入客户端
		s.connections++
		fmt.Println("已接入客户端")

		s.cond.L.Unlock()
	}
}

func main() {
	s := server{0,make(chan bool), sync.NewCond(&sync.Mutex{})}
	s.serve()
}
