package main

import "fmt"

// Dog can make a dog object
type Dog struct {
	name string
	age int
}

func (d *Dog) skill() {
	fmt.Println(d.name, "汪汪汪")
}

//type people struct {
//
//}

func newDog(name string, age int) *Dog {
	dog := Dog{}
	dog.name = name
	dog.age = age
	return &dog
}

func main() {
	//var dog1 = Dog{
	//	"臭屁",
	//	20,
	//}
	//
	//var dog *Dog = &Dog{
	//	name: "qiqi",
	//	age: 2,
	//}
	//
	//d := new(Dog)
	//d.name = "na"
	//d.age = 3
	//fmt.Println(d)
	//
	//fmt.Println(dog)
	//fmt.Println(dog1.name)
	//dog1.skill()
	dog := newDog("lili", 20)
	fmt.Println(dog)
	dog.skill()
}
