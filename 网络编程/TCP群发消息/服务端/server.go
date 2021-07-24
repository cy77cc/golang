package main

import (
	"fmt"
	"net"
	"os"
)

func checkError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:7111")
	checkError(err, "net.Listen")

	conns := make([]net.Conn, 0)
	for len(conns) < 3 {
		conn, err := listener.Accept()
		checkError(err, "listener.Accept")
		conns = append(conns, conn)
	}

	for _, conn := range conns {
		msg := fmt.Sprintf("Hello, %v", conn.RemoteAddr())
		_, err := conn.Write([]byte(msg))
		checkError(err, "conn.Write")
	}

}
