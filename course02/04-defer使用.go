package main

import "fmt"

func xingzuoZhengsuo() {
	fmt.Println("星座诊所")
}

func main() {
	fmt.Println("亲爱的患者，欢迎来到我院！")
	// 挂起一个延时任务（在当前函数“返回”前执行）
	defer fmt.Println("同志再见，我院永远欢迎您！")
	xingzuoZhengsuo()
	fmt.Println("事务A")
	fmt.Println("事务B")
	fmt.Println("事务C")
}
