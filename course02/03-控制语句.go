package main

import (
	"fmt"
	"time"
)

func main() {
	//for i := 1; i <= 9; i++ {
	//	for j := 1; j <= i; j++ {
	//		fmt.Printf("%d * %d = %d\t", i, j, i*j)
	//	}
	//	fmt.Println()
	//}
	var i int = 0
	loop: for {
		fmt.Printf("努力奋斗了%d小时\n", i)
		time.Sleep(500 * time.Millisecond)
		i++
		if i > 20 {
			goto out
		}
	}
	out:
		var choice string
		fmt.Println("外面的世界，要不要再玩一次Y/N")
		fmt.Scan(&choice)
		if choice == "Y" {
			goto loop
		}
}
