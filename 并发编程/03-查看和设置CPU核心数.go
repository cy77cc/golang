package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("当前协程的可用CPU逻辑核心数为：", runtime.NumCPU())

	// 把可用的最大逻辑CPU核心数设为4，返回先前的设置
	gomaxprocs := runtime.GOMAXPROCS(4)
	fmt.Println(gomaxprocs) // 8
	gomaxprocs = runtime.GOMAXPROCS(4)
	fmt.Println(gomaxprocs) // 4


}
