package main

import (
	"bufio"
	"fmt"
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
	listen, err := net.Listen("tcp", "127.0.0.1:7111")
	checkError(err, "net.Listen")
	defer func() {
		_ = listen.Close()
	}()
	conn, err := listen.Accept()
	checkError(err, "listen.Accept")
	defer func() {
		_ = conn.Close()
	}()
	file, err := os.OpenFile("D:\\learn\\golang\\网络编程\\文件上传作业\\服务端\\temp.iso", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	checkError(err, "os.OpenFile")
	defer func() {
		_ = file.Close()
	}()
	writer := bufio.NewWriter(file)
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		//fmt.Println(buf)
		checkError(err, "conn.Read")
		//_, err = file.Write(buf)
		_, err = writer.Write(buf[:n])
		checkError(err, "writer.Write")
		fmt.Println("文件接收中……")
		if n < 1024 {
			_, _ = writer.Write(buf[:n])
			_, _ = conn.Write([]byte("ok"))
			_ = writer.Flush()
			fmt.Println("文件写入成功")
			break
		}
	}

}
