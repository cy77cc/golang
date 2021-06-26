package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	/*
		随机生成一个3位数
		让用户输入其猜想
		反馈给用户：猜大了、猜小了、猜对了
		如果没猜对，就继续猜
		猜对了就退出游戏
		如果用户输入“exit"，就直接退出游戏
	*/
	myRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	var n = int64(myRand.Intn(1000))
	//fmt.Println(num)
	var num string
	for {
		fmt.Print("请输入一个数字：")
		fmt.Scan(&num)
		if num == "exit" {
			fmt.Println("游戏结束")
			break
		}
		num2, err := strconv.ParseInt(num, 0, 64)
		if err == nil {
			if num2 > n {
				fmt.Println("猜大了再来一遍！输入exit可以退出")
			} else if num2 < n {
				fmt.Println("猜小了再来一遍！输入exit可以退出")
			} else if num2 == n{
				fmt.Println("您猜对了")
				break
			}
		} else {
			fmt.Println("您输错了")
		}
	}
}
