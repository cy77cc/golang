package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, err := ioutil.ReadFile("E:\\golang\\文件读写\\file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
	//dir, err := ioutil.ReadDir("E:\\golang")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for _, v := range dir {
	//	fmt.Println(v.Name())
	//	fmt.Println(v.IsDir())
	//}
}
