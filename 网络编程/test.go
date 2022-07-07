package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var (
	reStr = `<div class="stem">(.+)</div>`
	//reStr = `<div class="right_answer"><span class="lit_title">答案</span>
	//                                        <div>
	//                                            <div>(.+)</div>
	//                                        </div>
	//                                    </div>`
	res = make([]string, 100)
)

func readfile(path string) string {
	file, _ := os.Open(path)
	bytes, _ := ioutil.ReadAll(file)
	defer func() {
		_ = file.Close()
	}()
	return string(bytes)
}

func writefile(path string, str string) {
	file, _ := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	defer func() {
		_ = file.Close()
	}()
	writer := bufio.NewWriter(file)
	_, _ = writer.Write([]byte(str))
	_ = writer.Flush()
}

func merge(A []string, B []string) {
	for i := 0; i < len(A); i++ {
		res[i] = fmt.Sprintf("**%s**\n\t%s\n", A[i], B[i])
	}
	for _, v := range res {
		//fmt.Println(v)
		writefile(`D:\learn\golang\网络编程\answer.md`, v)
	}
}

func main() {
	A := make([]string, 100)
	B := make([]string, 100)
	//str := readfile(`D:\learn\golang\网络编程\aaa.html`)
	//regUrl := regexp.MustCompile(reStr)
	//res := regUrl.FindAllStringSubmatch(str, -1)
	//i := 1
	//for _, v := range res {
	//	fmt.Println(v[1])
	//	//writefile(`D:\learn\golang\网络编程\quest.txt`, fmt.Sprintf("%d.%s\n", i, v[1]))
	//	i++
	//}
	file, _ := os.Open(`D:\learn\golang\网络编程\quest.txt`)
	file2, _ := os.Open(`D:\learn\golang\网络编程\answer.txt`)
	defer func() {
		_ = file.Close()
		_ = file2.Close()
	}()
	reader := bufio.NewReader(file)
	for i := 0; ; i++ {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		A[i] = string(line)
	}
	reader = bufio.NewReader(file2)

	for i := 0; ; i++ {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		B[i] = string(line)
	}
	merge(A, B)
}
