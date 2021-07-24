package main

import (
	"fmt"
	"net"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	// 解析UDP地址
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:6333")
	checkError(err)

	// 服务器建立监听
	serverConn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	fmt.Println("正在监听6333端口")
	defer func() {
		_ = serverConn.Close()
	}()

	// 与客户端IO
	var content string
	buf := make([]byte, 1024)
	for {
		//读取客户端消息到缓冲区，
		//返回值 n 是字节数，
		//remoteAddress 数据来源的地址，
		//err 错误
		n, remoteAddress, _ := serverConn.ReadFromUDP(buf)
		contents := string(buf[:n])
		fmt.Printf("来自%v:%v的消息: %v\n", udpAddr.IP, udpAddr.Port, contents)
		if n != 0 {
			fmt.Print("输入聊天内容(输入exit退出)：")
			_, _ = fmt.Scan(&content)
			if content == "exit" {
				return
			}
			// 服务端如果直接使用write的话他会不知道发给谁，然后报错，所以在读的时候使用ReadFromUDP，可以获取一个地址
			_, _ = serverConn.WriteToUDP([]byte(content), remoteAddress)
		}
	}
}
