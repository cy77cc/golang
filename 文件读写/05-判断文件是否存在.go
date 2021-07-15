package main

import (
	"fmt"
	"os"
)

func main() {
	stat, err := os.Stat("E:\\golang\\文件读写\\file.txt")
	if stat != nil && err == nil {
		fmt.Println("文件存在")
	} else if os.IsNotExist(err) {
		fmt.Println("文件不存在")
	} else {
		fmt.Println(err)
	}
}
