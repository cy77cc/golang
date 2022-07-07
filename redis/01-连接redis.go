package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
)

func handleError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}

func main() {
	// 连接redis数据库
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	handleError(err, "redis.Dial")

	defer func() {
		_ = conn.Close()
	}()

	// 执行redis命令，获得结果
	reply, err := conn.Do("set", "name", "张三")
	handleError(err, "conn.Do set")
	fmt.Printf("type=%T, value=%v", reply, reply)
	reply, err = conn.Do("get", "name")
	handleError(err, "conn.Do set")
	fmt.Printf("type=%T, value=%s", reply, reply)
}