package main

import (
	"fmt"
	"net"
	"strings"
)

/*
负责注册节点名称：节点运行端口
*/

var peerNameAddressMap map[string]string

func main() {
	//初始化注册表
	peerNameAddressMap = make(map[string]string)
	//开启注册服务
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
		return
	}

	buffer := make([]byte, 1024)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}
		//读取节点名称
		n, _ := conn.Read(buffer)
		message := string(buffer[:n])

		strs := strings.Split(message, " ")
		cmd := strs[0]
		peerName := strs[1]
		if cmd == "reg" {
			//讲节点名称和节点地址写入映射表
			address := conn.RemoteAddr().String()

			//listenAddress := strings.Split(address, ":")[0]+ strs[2]
			address = strings.Split(address, ":")[0] + ":" + strs[2]
			peerNameAddressMap[peerName] = address

			//msg := fmt.Sprintf("%s-%s注册成功", peerName, address)
			_, _ = conn.Write([]byte(address))
		} else if cmd == "get" {
			address := peerNameAddressMap[peerName]
			_, _ = conn.Write([]byte(address))
		} else {
			fmt.Println("参数不存在")
		}
		_ = conn.Close()
	}
}
