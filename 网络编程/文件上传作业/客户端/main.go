package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func checkError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		fmt.Println("客户端异常退出")
		// 退出
		os.Exit(1)
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:7111")
	checkError(err, "net.Dial")
	defer func() {
		_ = conn.Close()
		fmt.Println("客户端正常退出")
	}()
	file, err := os.Open(`D:\Centos 7\CentOS-7-x86_64-DVD-1908.iso`)
	checkError(err, "os.Open")
	defer func() {
		_ = file.Close()
	}()
	buf := make([]byte, 1024)
	reader := bufio.NewReader(file)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			fmt.Println("文件上传成功")
			break
		}
		_, _ = conn.Write(buf[:n])
		fmt.Println("文件发送中……")
	}
	n, _ := conn.Read(buf)
	fmt.Println(string(buf[:n]))
}
