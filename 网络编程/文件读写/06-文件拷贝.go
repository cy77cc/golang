package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	/*
		io.Copy 文件拷贝API
	*/
	file1, err := os.OpenFile("E:\\golang\\文件读写\\file.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		file2, err2 := os.OpenFile("E:\\golang\\文件读写\\file2.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
		if err2 != nil {
			fmt.Println(err2)
			return
		} else {
			reader := bufio.NewReader(file1)
			writer := bufio.NewWriter(file2)
			for {
				readString, err3 := reader.ReadString('\n')
				_, _ = writer.WriteString(readString)
				if err3 == io.EOF {
					break
				}
			}
			_ = writer.Flush()
		}
		defer func() {
			_ = file2.Close()
		}()
	}
	defer func() {
		_ = file1.Close()
	}()
}
