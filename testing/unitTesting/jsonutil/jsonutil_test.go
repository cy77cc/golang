package jsonutil

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncodeHuman2File(t *testing.T) {
	// 虚构一个Human数据
	h := Human{"张三", 20, true}
	filename := "E:\\golang\\testing\\unitTesting\\jsonutil\\test.json"
	// 调用编码方法将数据编码到json文件
	err := EncodeHuman2File(&h, filename)
	if err != nil {
		t.Fatal("出错了", err)
	}
	//	将编码得到的文件进行解码
	h2 := new(Human)
	file, err := os.Open(filename)
	if err != nil {
		t.Fatal("文件写入失败", err)
	}
	defer func() {
		_ = file.Close()
	}()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(h2)
	if err != nil {
		t.Fatal("解码失败", err)
	}
	//与原始数据做对比
	if h.Name == h2.Name && h.Age == h2.Age && h.Gender == h2.Gender {
		t.Log("测试成功")
	} else {
		t.Fatal("编码失败")
	}
}

func TestDecodeJsonFile(t *testing.T) {
	// 虚构一个Human数据
	h := Human{"张三", 20, true}
	filename := "E:\\golang\\testing\\unitTesting\\jsonutil\\test.json"
	//	将编码得到的文件进行解码
	h2 := new(Human)
	err := DecodeJsonFile(filename, h2)
	if err != nil {
		t.Fatal("解码失败", err)
	}
	//与原始数据做对比
	if h.Name == h2.Name && h.Age == h2.Age && h.Gender == h2.Gender {
		t.Log("测试成功")
	} else {
		t.Fatal("编码失败")
	}
}
