package main

import (
	"context"
	"fmt"
	"time"
)

/*func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")

	cancel()

	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}*/

func main() {
	ctx, cancel1 := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithCancel(context.Background())

	go func(ctx context.Context, cancel2 context.CancelFunc) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("协程1退出了")
				fmt.Println("协程2你也给我退出吧")
				cancel2()
				return
			default:
				fmt.Println("协程执行中")
				time.Sleep(time.Second)
			}
		}
	}(ctx, cancel2)

	go func(ctx2 context.Context, cancel context.CancelFunc) {
		time.Sleep(5 * time.Second)
		fmt.Println("协程1要被结束了")
		cancel()
		for {
			select {
			case <-ctx2.Done():
				fmt.Println("协程2也结束了")
				return
			default:

			}
		}
	}(ctx2, cancel1)

	time.Sleep(10 * time.Second)
}
