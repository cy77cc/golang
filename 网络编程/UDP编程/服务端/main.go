package main

import (
	"fmt"
	"net"
)

func main() {
	// 解析udp地址
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:6100")
	if err != nil {
		fmt.Println(err)
		return
	}
	//监听udp
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, _ = conn.Write([]byte("Hello"))
	defer func() {
		_ = conn.Close()
	}()

}
