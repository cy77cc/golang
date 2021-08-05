package main

import (
	"fmt"
	"sync"
	"time"
)

/*
城管预警
监听城管大队
烧烤摊集群：监听城管大队，只要出动就发消息通知工会主席进入阻塞等待至唤醒，否则就提供露天烧烤
工会主席：摆平城管大队，并广播通知所有烧烤摊主
*/

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	var dangerous = false
	dangerousMsg := make(chan string, 1)

	for i := 0; i < 3; i++ {
		go func(n int) {
			for {
				//只要城管出来，就等待起来
				cond.L.Lock()
				for dangerous {
					//发送预警
					select {
					case dangerousMsg <- "城管来了！！！！":
						fmt.Println("城管来了！！！！")
						fmt.Println("进入蛰伏状态")
					default:

					}
					cond.Wait()
				}
				cond.L.Unlock()

				//城管没有出来
				if n == 0 {
					fmt.Println(n+1, "露天烧烤。。。")
				} else if n == 1 {
					fmt.Println(n+1, "麻辣烫。。。。")
				} else {
					fmt.Println(n+1, "武大郎烧饼。。。。")
				}
				time.Sleep(2*time.Second)
			}
		}(i)
	}


	//工会主席
	go func() {
		for {
			select {
			//城管来了，平事
			case <-dangerousMsg:
				cond.L.Lock()
				fmt.Println("主席出动。。。。")
				time.Sleep(2*time.Second)
				dangerous = false
				fmt.Println("事情已经摆平了。。。。")
				cond.Broadcast()
				cond.L.Unlock()
			//	日常生活
			default:
				fmt.Println("工会主席摸鱼。。。。")
				time.Sleep(4 * time.Second)
				dangerous = true
			}
		}
	}()

	for {
		fmt.Print()
	}
}
