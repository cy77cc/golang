package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	err := ioutil.WriteFile("E:\\golang\\文件读写\\file.txt", []byte("太棒了"), 0666)
	if err != nil {
		fmt.Println("写入文件失败", err)
		return
	} else {
		fmt.Println("写入成功")
	}
}
