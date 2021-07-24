package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6444")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		_ = conn.Close()
	}()
	_, err = conn.Write([]byte("Hello"))
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Print(string(buf))
}
