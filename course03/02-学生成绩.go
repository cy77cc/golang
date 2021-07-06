package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sort(score map[string]int) {
	slice := make([]int, 0)
	var tempMap = map[int]string{}
	for key, value := range score {
		slice = append(slice, value)
		tempMap[value] = key
	}
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] < slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	fmt.Println(slice)
	for i := 0; i < len(slice); i++ {
		fmt.Printf("第%d名：%s, 成绩：%d\n", i+1, tempMap[slice[i]], slice[i])
	}
}

func main() {
	score := map[string]int{}
	for {
		var name string
		fmt.Print("输入学生姓名（输入exit退出）：")
		_, _ = fmt.Scan(&name)
		if name == "exit" {
			break
		}
		var num int
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		num = r.Intn(101)
		score[name] = num
		time.Sleep(10 * time.Nanosecond)
	}
	fmt.Println(score)
	sort(score)
	fmt.Println(score)
}
