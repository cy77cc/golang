package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type Good struct {
	Name  string
	Price float64
}

func newGood(name string, price float64) string {
	good := Good{name, price}
	bytes, _ := json.Marshal(good)
	return string(bytes)
}

func main01() {
	goods := make([]string, 5)
	goods = append(goods, newGood("书本", 30))
	goods = append(goods, newGood("水果", 40))
	goods = append(goods, newGood("牛奶", 30))
	goods = append(goods, newGood("洗衣粉", 60))
	goods = append(goods, newGood("牙膏", 30))
	file, err := os.OpenFile(`E:\golang\SDK\test.csv`, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		_ = file.Close()
	}()
	writer := csv.NewWriter(file)
	_ = writer.Write(goods)
	writer.Flush()

}


func main() {
	file, err := os.Open(`C:\Users\48356\Desktop\录取.CSV`)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		_ = file.Close()
	}()

	reader := csv.NewReader(file)
	all, _ := reader.ReadAll()
	fmt.Println(all)
}