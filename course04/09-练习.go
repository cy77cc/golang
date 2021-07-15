package main

import "fmt"

type IPerson interface {
	eat()
	drink()
	sleep()
}

type Person struct {
	name string
}

func (p *Person) eat() {
	fmt.Println("吃")
}

func (p *Person) drink() {
	fmt.Println("喝")
}

func (p *Person) sleep() {
	fmt.Println("睡觉")
}

type IWorker interface {
	// output 是工作产出
	IPerson
	work() (output string)
	rest()
}

type worker struct {
	Person
	skill  string
	hours  int
	title  string
	output string
}

func (w *worker) work() (output string) {
	fmt.Printf("%s正在%s\n", w.name, w.skill)
	return w.output
}

func (w *worker) rest() {
	fmt.Printf("%s正在休息\n", w.name)
}

type coder struct {
	worker
}

type teacher struct {
	worker
}

type farmer struct {
	worker
}

func main() {
	workers := make([]IWorker, 0)
	workers = append(workers, &coder{
		worker{
			Person{"coder"},
			"敲代码",
			12,
			"高级工程师",
			"代码",
		}})
	workers = append(workers, &teacher{
		worker{
			Person{
				"teacher",
			},
			"教书",
			8,
			"特级教师",
			"知识",
		},
	})
	workers = append(workers, &farmer{
		worker{
			Person{
				"farmer",
			},
			"种田",
			10,
			"农民",
			"粮食",
		},
	})

	for i := 1; i <= 7; i++ {
		if (i == 6) || (i == 7) {
			for _, v := range workers {
				if _, ok := v.(*coder); ok {
					v.work()
				} else if _, ok := v.(*teacher); ok {
					v.rest()
				} else {
					v.eat()
					v.drink()
					v.sleep()
				}
			}
		} else {
			for _, v := range workers {
				v.work()
			}
		}
	}

}
