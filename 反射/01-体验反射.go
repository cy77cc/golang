package main

import (
	"fmt"
	"reflect"
)

type Human struct {
	Name string `json:"name"`
}

func (h Human) SayName() {
	fmt.Println(h.Name)
}

func (h Human) SayHello() {
	fmt.Println("Hello")
}

func (h *Human) SayYes() {
	fmt.Println("yes")
}

func (h *Human) SayYY(name string, age int) {
	fmt.Println(name, age)
}

func main() {
	//human := Human{"张三"}
	//返回human的类型
	//typeof := reflect.TypeOf(human)
	//fmt.Println(typeof) // main.Human

	//返回typeof的系统类别
	//fmt.Println(typeof.Kind()) // struct

	//获得typeof属性的个数
	//fmt.Println(typeof.NumField()) // 1

	//获得typeof方法的个数
	//fmt.Println(typeof.NumMethod()) // 1

	//获得第i个属性，越界会报错
	//field := typeof.Field(0)

	//获取第一个属性的字段名
	//fmt.Println(field.Name)

	//获取第一个属性的类型
	//fmt.Println(field.Type)

	//获取第一个属性的标签
	//fmt.Println(field.Tag)

	//获取第一个方法
	//method := typeof.Method(0)

	//获取第一个方法的名字
	//fmt.Println(method.Name)

	//获取第一个方法的类型
	//fmt.Println(method.Type)

	//获取第一个方法的包路径
	//fmt.Println(method.PkgPath)

	//找出第0个属性里的第1个属性
	//typeof.FieldByIndex([]int{0, 1})

	/*--------------------------------------------------------------------------------*/

	//valueOf := reflect.ValueOf(human)
	//fmt.Println(valueOf)

	//对象第0个属性的值
	//vField := valueOf.Field(0)
	//反射值转正射值
	//fmt.Println(valueOf.Interface())
	//fmt.Println(vField.Interface())

	//对象第0个属性里的第1个属性的值
	//valueOf.FieldByIndex([]int{0, 1})

	h2 := &Human{"李四"}
	ptrValue := reflect.ValueOf(h2)
	ptrType := reflect.TypeOf(h2)
	method := ptrType.Method(2)
	fmt.Println(method.Name)

	//获取method的参数个数（包括了结构体自己），NumOut是返回值个数
	fmt.Println(method.Type.NumIn())

	//获取第i个参数类型
	fmt.Println(method.Type.In(1))

	//获取地址或者接口值得Value封装，反射值
	//elem := ptrValue.Elem()
	//fmt.Println(elem)
	//fmt.Println(elem.Interface())

	//获取ptrValue的方法个数，指针类型可以获取所有的方法，指针可以替代值，值不可以替代指针
	//fmt.Println(ptrValue.NumMethod())

	//获得当前value的第i个的值,内存中的代码区里的一堆指令
	methodValue := ptrValue.Method(2)
	fmt.Println(methodValue)
	//等同于h2.SayHello("中华田园犬")，有参数的写参数，没参数空切片
	methodValue.Call([]reflect.Value{
		reflect.ValueOf("中华田园犬"),
		reflect.ValueOf(20),
	})

	//h3 := Human{"赵五"}
	//h3Value := reflect.ValueOf(h3)
	//nameField := h3Value.Field(0)
	//检查当前地址value内的值是否可以改变
	//可改变条件：可寻址+不来自非导出字段
	//fmt.Println(nameField.CanSet())

	//获取h3Value的方法个数，只可以获取值绑定的方法个数
	//fmt.Println(h3Value.NumMethod())

	//var a int
	//aValue := reflect.ValueOf(&a).Elem()
	//必须使用指针类型
	//aValue.SetInt(20)
	//fmt.Println(a)
}
