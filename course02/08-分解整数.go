package main

import (
	"fmt"
	_ "fmt"
)

func resolveInt(num int) (hunderd, decade, bit int) {
	hunderd = num / 100
	decade = num % 100 / 10
	bit = num % 10
	return
}

func main() {
	hunderd, decade, bit := resolveInt(654)
	fmt.Println(hunderd, decade, bit)
}
