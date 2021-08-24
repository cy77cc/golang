package main

import (
	"fmt"
	"os"
)

func main02() {
	// 返回当前工作目录的根目录
	getwd, err := os.Getwd()
	if err == nil {
		fmt.Println(getwd)
	} else {
		fmt.Println(err)
	}

	// 获取环境变量
	//getenv := os.Getenv("JAVA_HOME")
	//fmt.Println(getenv)

	// 获得全部环境变量
	//environ := os.Environ()
	//fmt.Println(environ)

	// 打印当前设备在网络中的主机名
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println(err)
	}

	// 获取当前用户的临时文件夹位置
	dir := os.TempDir()
	fmt.Println(dir)

	// 判断当前系统的文件分隔符
	separator := os.IsPathSeparator('\\')
	fmt.Println(separator)

	stat, err := os.Stat("E:\\golang\\SDK\\01-strings包.go")
	if err == nil {
		fmt.Println(stat.IsDir())
		fmt.Println(stat.Mode())
		fmt.Println(stat.Name())
		fmt.Println(stat.Size())
		fmt.Println(stat.ModTime())
	} else {
		fmt.Println(err)
	}
}

func main() {
	file, err := os.OpenFile(`E:\golang`, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("openfile", err)
		return
	}

	defer func() {
		_ = file.Close()
	}()

	readDir, _ := file.ReadDir(-1)
	for i := 0; i < len(readDir); i++ {
		fmt.Println(readDir[i].Name(), readDir[i].IsDir())
	}
}