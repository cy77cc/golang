package main

import (
	"fmt"
	"os"
)

//const dir string = `E:\golang\SDK\out\`

func main() {
	//file, _ := os.OpenFile(dir+"test.gz", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	//defer func() {
	//	_ = file.Close()
	//}()
	//
	//writer := gzip.NewWriter(file)
	//
	//_ = writer.Flush()

	file, _ := os.Open(`E:\golang\SDK`)
	dir, _ := file.ReadDir(-1)
	for _, v := range dir {
		fmt.Println(v.Name())
		fmt.Println(v.Info())
	}

}
