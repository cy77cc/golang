package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:6333")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		_ = conn.Close()
	}()

	// 客户端发起交谈
	conn.Write([]byte("Hello"))
	buf := make([]byte, 1024)
	var content string
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		contents := string(buf[:n])
		fmt.Printf("来自的消息: %v\n", contents)
		if n != 0 {
			fmt.Print("输入聊天内容(输入exit退出)：")
			_, _ = fmt.Scan(&content)
			if content == "exit" {
				return
			}
			_, _ = conn.Write([]byte(content))
		}
	}
}
