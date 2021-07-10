package main

import "fmt"

type USB interface {
	start() bool
	stop() bool
}

type phone struct {
	name string
}

func (p *phone) start() bool {
	fmt.Println(p.name + "启动了")
	return true
}

func (p *phone) stop() bool {
	fmt.Println(p.name + "关闭了")
	return true
}

func main() {
	// 用子类实现去为接口赋值
	var usb USB
	p1 := phone{"iphone4"}
	// 这里之所以要赋值地址是因为，在实现接口方法时是用的指针类型，所以要赋值指针类型
	// 假如实现方法绑定不是指针类型则可以直接赋值，不用赋值地址
	usb = &p1
	usb.start()
	usb.stop()

}
