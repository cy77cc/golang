package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 查重
type rg struct {
	id       int
	question string
	answer   string
}

var n, m int
var res2 = make([]rg, 100)
var norepeat = make([]rg, 100)

func readfileline(path string) {
	path1 := path + `\quest.txt`
	path2 := path + `\answer.txt`
	file1, _ := os.Open(path1)
	file2, _ := os.Open(path2)
	defer func() {
		_ = file1.Close()
		_ = file2.Close()
	}()
	reader1 := bufio.NewReader(file1)
	reader2 := bufio.NewReader(file2)
	for i := 0; ; i++ {
		bytes1, _, err1 := reader1.ReadLine()
		bytes2, _, _ := reader2.ReadLine()
		if err1 == io.EOF {
			break
		}
		res2[i] = rg{i + 1, string(bytes1), string(bytes2)}
		n = i + 1
	}
}

func deleteRepeat() {
	norepeat[0] = res2[0]
	m++
	for i := 1; i < n; i++ {
		flag := true
		for j := 0; j < m; j++ {
			if res2[i].question == norepeat[j].question {
				flag = false
			}
		}
		if flag {
			norepeat[m] = res2[i]
			m++
		}
	}
}

func main() {
	readfileline(`D:\learn\golang\网络编程`)
	//fmt.Println(n)
	deleteRepeat()
	file, _ := os.OpenFile(`D:\learn\golang\网络编程\answer2.md`, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer func() {
		_ = file.Close()
	}()
	writer := bufio.NewWriter(file)
	for i := 0; i < m; i++ {
		_, _ = writer.WriteString(fmt.Sprintf("#### %d.%s\n\t%s\n", i+1, norepeat[i].question, norepeat[i].answer))
	}
	_ = writer.Flush()
}
