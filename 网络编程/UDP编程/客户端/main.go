package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:6100")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		_ = conn.Close()
	}()
	_, err = conn.Write([]byte("你好呀"))
	if err != nil {
		fmt.Println(err)
		return
	}
	var buf []byte
	buf = make([]byte, 0)
	_, _ = conn.Read(buf)
	fmt.Println(buf)
}
