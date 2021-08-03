package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)

	var count int
	for {
		<-ticker.C
		fmt.Println("我要去浪")
		count++
		if count >= 9 {
			ticker.Stop()
			break
		}
	}
}
