package main

import "fmt"

var Name = "lillli"

type wildAnimal interface {
	// 抽象方法
	born()
	die()
}

type Tiger interface {
	// 显示继承
	wildAnimal
	kill()
}

type leopard interface {
	// 隐式继承
	born()
	die()
	run()
}

func Say() {
	fmt.Println("Say")
}

//func main() {
//
//}
