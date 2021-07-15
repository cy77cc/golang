package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开文件，给读写权限
	file, err := os.OpenFile("E:\\golang\\文件读写\\file.txt", os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		// 创建io.Reader对象
		//reader := io.Reader(file)

		// 创建io.Writer对象
		//writer := io.Writer(file)

		// 创建带缓冲的读对象，参数是io.Reader对象
		// file实现了io.Reader/io.Writer接口
		newReader := bufio.NewReader(file)

		// 创建带缓冲的写对象，参数是io.Writer对象
		newWriter := bufio.NewWriter(file)
		for {
			// 用传入的分隔符作为每一次读的结束
			readString, err := newReader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				// io.EOF 就是读到文件末尾的错误
				if err == io.EOF {
					break
				}
			}
			fmt.Print(readString)
		}
		for i := 0; i < 10; i++ {
			// 带缓冲的写入操作
			_, err := newWriter.Write([]byte("\n你好"))
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		// 带缓冲的写入操作在结束时要使用，强制将缓冲区中的内容写入文件
		_ = newWriter.Flush()
	}

	defer func() {
		err2 := file.Close()
		if err2 != nil {
			fmt.Println(err2)
		}
	}()
}
