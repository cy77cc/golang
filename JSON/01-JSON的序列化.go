package main

import (
	"encoding/json"
	"fmt"
)

// 要使用json就必须大写，否则json包访问不到
type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender bool   `json:"gender"`
}

func main() {
	person := Person{"张三", 27, true}
	marshal, err := json.Marshal(person)
	if err != nil {
		fmt.Println("转换失败", err)
		return
	} else {
		fmt.Println(string(marshal))
	}

	var p2 Person
	err = json.Unmarshal(marshal, &p2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(p2)
	}

}
