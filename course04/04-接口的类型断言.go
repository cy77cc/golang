package main

import "fmt"

type animal interface {
	born()
	die()
}

type dog struct {
	name string
}

func (d *dog) born() {
	fmt.Println(d.name + "出生了")
}

func (d *dog) die() {
	fmt.Println(d.name + "挂了")
}

type cat struct {
	name string
}

func (c *cat) born() {
	fmt.Println(c.name + "出生了")
}

func (c *cat) die() {
	fmt.Println(c.name + "挂了")
}

func main() {
	animals := make([]animal, 2)
	animals[0] = &dog{"旺财"}
	animals[1] = &cat{"咪咪"}
	for i := 0; i < len(animals); i++ {
		switch animals[i].(type) {
		case *dog:
			fmt.Println("dog")
		case *cat:
			fmt.Println("cat")
		default:
			fmt.Println("不知道")
		}
	}
}
