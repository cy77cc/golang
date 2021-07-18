package main

import (
	"fmt"
)

func main() {
	defer func() {
		err := recover()
		//fmt.Println(err)
		s, ok := err.(string)
		if ok && s == "你妹"{
			fmt.Println("你妹妹了")
		}
	}()
	panic("你妹")
}
