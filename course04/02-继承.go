package main

import "fmt"

// Person a person class
type Person struct {
	name string
	age int
	gender bool
}

// Student Student inherit from Person and can make student object
type Student struct {
	// 继承父类的所有属性和方法
	Person
	grade int
}

//skill Person method
func (p *Person) skill() {
	fmt.Println("person学习", p.name)
	p.name = "李四"
	fmt.Println("person学习", p.name)
}

// 重写父类的方法
func (s *Student) skill() {
	fmt.Printf("student%s学习考了%d", s.name, s.grade)
}

func main() {
	person := Person{
		name: "张三",
		age: 20,
		gender: true,
	}
	person.skill()
	fmt.Println(person.name)
	//student := Student{}
	//student.name = "小明"
	//student.age = 12
	//student.gender = true
	//student.grade = 89
	//student.skill()

	xm := Student{Person{name: "小明", age: 20, gender: true}, 90}
	xm.skill()
}
