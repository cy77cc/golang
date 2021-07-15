package main

import "fmt"

// Person1 a Person1 class
type Person1 struct {
	name string
	age int
	gender bool
}

// Student Student inherit from Person1 and can make student object
type Student struct {
	// 继承父类的所有属性和方法
	Person1
	grade int
}

//skill Person1 method
func (p *Person1) skill() {
	fmt.Println("Person1学习", p.name)
	p.name = "李四"
	fmt.Println("Person1学习", p.name)
}

// 重写父类的方法
func (s *Student) skill() {
	fmt.Printf("student%s学习考了%d", s.name, s.grade)
}

func main() {
	person1 := Person1{
		name: "张三",
		age: 20,
		gender: true,
	}
	person1.skill()
	fmt.Println(person1.name)
	//student := Student{}
	//student.name = "小明"
	//student.age = 12
	//student.gender = true
	//student.grade = 89
	//student.skill()

	xm := Student{Person1{name: "小明", age: 20, gender: true}, 90}
	xm.skill()
}
