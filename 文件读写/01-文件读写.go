package main

import (
	"fmt"
	"os"
)

/*
	buffer 缓冲区
	utility 便利的工具
	util 便捷工具
*/

func main() {
	//file, err := os.Open("E:\\golang\\文件读写\\file.txt")
	//if err == nil {
	//	fmt.Println(file.Name())
	//	return
	//} else{
	//	fmt.Println("文件打开错误", err)
	//}
	//defer func () {
	//	err2 := file.Close()
	//	if err2 != nil {
	//		fmt.Println(err2)
	//	}
	//}()

	file, err := os.OpenFile("E:\\golang\\文件读写\\file2.txt",
		os.O_RDWR|os.O_CREATE, 0666)
	if err == nil {
		fmt.Println(file.Name())
	} else {
		fmt.Println(err)
	}

	defer func() {
		err2 := file.Close()
		if err2 != nil {
			fmt.Println(err2)
		}
	}()
}
