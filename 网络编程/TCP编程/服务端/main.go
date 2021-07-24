package main

import (
	"fmt"
	"net"
)

func main() {
	//建立监听
	listen, err := net.Listen("tcp", "127.0.0.1:6444")
	if err != nil {
		fmt.Println(err)
		return
	}
	//接入一个客户端
	conn, err := listen.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		_ = conn.Close()
	}()
	fmt.Println("成功接入客户端", conn.RemoteAddr().String())
	//在专线连接中与特定的客户端IO
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Println(string(buf[:n]))
	_, _ = conn.Write([]byte("Hello"))
}
