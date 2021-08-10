package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
开辟三条协程：
A每秒生成一个三位随机数
B输出该值的奇偶性
C输出奇数和偶数的总个数
当生成一个水仙花数时，程序结束
*/

func main() {

	numCh := make(chan int, 2)
	done := make(chan int)
	oddCh := make(chan int, 1)
	evenCh := make(chan int, 1)
	var oddTotal int
	var evenTotal int

	//A每秒生成一个三位随机数
	go func() {
		for {
			ticker := time.NewTicker(100*time.Millisecond)
			<-ticker.C
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			i := r.Intn(900) + 100
			fmt.Println("生成", i)
			numCh <- i
		}
	}()

	//B输出该值的奇偶性
	go func() {
		for {
			v := <-numCh
			if v%2 == 0 {
				evenCh <- v
				fmt.Println(v, "是偶数")
			} else {
				oddCh <- v
				fmt.Println(v, "是奇数")
			}

		}
	}()

	//C输出奇数和偶数的总个数
	go func() {
		var v int
		for {
			select {
			case v = <-evenCh:
				evenTotal++
				a := v / 100
				b := v / 10 % 10
				c := v % 10

				if a*a*a+b*b*b+c*c*c == v {
					done <- v
				}
			case v = <-oddCh:
				oddTotal++
				a := v / 100
				b := v / 10 % 10
				c := v % 10

				if a*a*a+b*b*b+c*c*c == v {
					done <- v
				}
			}
		}
	}()

	v := <-done
	fmt.Println(v)
	fmt.Println("奇数有：", oddTotal)
	fmt.Println("偶数有：", evenTotal)

}
