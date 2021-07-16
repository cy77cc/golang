package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Human struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Male bool   `json:"male"`
}

func main() {
	file, err := os.OpenFile("E:\\golang\\JSON\\data.json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		_ = file.Close()
	}()

	human := Human{"李四", 20, false}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(human)
	if err != nil {
		fmt.Println(err)
	}

	file2, err := os.Open("E:\\golang\\JSON\\data.json")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	defer func() {
		_ = file2.Close()
	}()

	human2 := Human{}
	// 解码JSON文件为go数据
	decoder := json.NewDecoder(file2)
	err = decoder.Decode(&human2)
	if err != nil {
		fmt.Println("err =", err)
	} else {
		fmt.Println(human2)
	}
}
