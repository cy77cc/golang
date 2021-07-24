package main

import (
	"net"
	"fmt"
	"os"
)

func checkError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:7111")
	checkError(err, "net.Dial")
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	checkError(err, "conn.Read")
	content := string(buffer[:n])
	fmt.Println(content)
}
