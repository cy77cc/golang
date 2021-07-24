package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

/*
用一个可执行程序实现相互聊天
实现注册节点名称，并通过名称发起会话
实现群发消息
*/
var peerPort, peerName *string

func StartServe() {

	listener, err := net.Listen("tcp", ":"+*peerPort)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}
		//在一个独立的并发任务中与其他节点IO
		go func(conn net.Conn) {
			buffer := make([]byte, 1024)
			n, _ := conn.Read(buffer)
			msg := string(buffer[:n])
			fmt.Printf("来自%v的消息: %s\n", conn.RemoteAddr(), msg)
			_ = conn.Close()
		}(conn)
	}
}

func StartRequest() {
	/*能主动地与其他节点发起会话*/
	var targetName, msg string
	for {
		fmt.Println("请输入对方名称：消息内容")
		_, _ = fmt.Scan(&targetName, &msg)

		//向注册器索取节点的地址
		targetAddress := RegGetPeerAddress("get " + targetName)
		conn, err := net.Dial("tcp", targetAddress)
		if err != nil {
			fmt.Println(err)
			break
		}
		_, _ = conn.Write([]byte(msg))
		_ = conn.Close()
	}

}

func RegGetPeerAddress(req string) string {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
	}
	_, _ = conn.Write([]byte(req))
	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer)
	return string(buffer[:n])
}

func main() {
	peerNameAddressMap = make(map[string]string)
	// 从命令行接收一个用于监听的端口
	peerPort = flag.String("port", "7111", "port")
	peerName = flag.String("name", "zhang", "peername")
	flag.Parse()

	//注册服务
	peerAddress := RegGetPeerAddress("reg " + *peerName + " " + *peerPort)
	fmt.Println("节点注册成功", *peerName, peerAddress)

	//接受聊天请求, 能够被动地被节点攀谈
	go StartServe()
	go StartRequest()

	// 阻塞主协程
	for {
		time.Sleep(time.Second * 1)
	}
}
