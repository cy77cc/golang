package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Year()) // 2021
	//fmt.Println(now.YearDay())
	fmt.Println(now.Weekday()) // Thursday
	fmt.Println(now.Month()) // July
	fmt.Println(now.Day()) // 8
	fmt.Println(now.Date()) // 2021 July 8
	fmt.Println(now.Hour()) // 10
	fmt.Println(now.Minute()) // 3
	fmt.Println(now.Second()) // 17
	fmt.Println(now.String())

	date := time.Date(2020, 2, 20, 20, 20, 13, 9, time.Now().Location())
	fmt.Println(date)

	// 两个时间相减
	fmt.Println(now.Sub(date))

	// 格式化时间差
	duration, _ := time.ParseDuration("12086h5m50.550331391s")
	fmt.Println(date.Add(duration))
}
