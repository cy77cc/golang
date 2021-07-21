package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p *Person) SayName() {
	fmt.Println(p.Name)
}

func (p Person) SayAge() {
	fmt.Println(p.Age)
}

func main() {
	obj := Person{"张三", 27}
	// 获取的其类型和值
	oType := reflect.TypeOf(obj)
	oValue := reflect.ValueOf(obj)
	fmt.Println(oType)
	fmt.Println(oType.Kind())
	fmt.Println(oValue)
	// 获得所有属性
	numField := oType.NumField()
	for i := 0; i < numField; i++ {
		//拿到所有属性的类型和名称
		fmt.Println(oType.Field(i))

		//获取所有属性的值
		fmt.Println(oValue.Field(i).Interface())
	}

	// 获得所有方法类型和名值
	oPtrType := reflect.TypeOf(&obj)
	numMethod := oPtrType.NumMethod()
	fmt.Println(numMethod)
	for i := 0; i < numMethod; i++ {
		method := oPtrType.Method(i)
		fmt.Println(method.Name, method.Type)

		//获取方法参数列表
		//获得参数个数
		numArgs := method.Type.NumIn()
		for j := 0; j < numArgs; j++ {
			//获得第j个参数
			argType := method.Type.In(j)
			fmt.Println(argType, argType.Name(), argType.Kind())
		}
	}

	// 访问属性
	oPtrValue := reflect.ValueOf(&obj)
	oPtrValue.Elem().Field(0).SetString("李四")
	oPtrValue.Elem().Field(1).SetInt(30)
	fmt.Println(obj)

	// 调用方法7
	method := oPtrValue.Method(0)
	method.Call([]reflect.Value{})
}