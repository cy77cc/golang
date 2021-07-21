package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

/*
所有商品都有一些共性，例如都有品名，价格，个性则无千无万
自行封装出三中商品（以模拟30万种商品）
随意给出一个商品的集合，将每件商品的所有属性值输出到《品名.txt》文件中
*/

type Computer struct {
	Name   string
	Price  float64
	Cpu    string
	Memory int
	Disk   int
}

type TShirt struct {
	Name  string
	Price float64
	Color string
	Size  int
	Sex   bool
}

type Car struct {
	Name  string
	Price float64
	Cap   int
	Power string
}

func main() {
	goods := make([]interface{}, 0)
	goods = append(goods, Computer{"未来人类", 10000, "Core i7", 16, 1024})
	goods = append(goods, TShirt{"爆款T恤", 10000, "红色", 40, false})
	goods = append(goods, Car{"兰博基尼", 100000, 5, "油电混合"})
	for _, v := range goods {
		//通过反射拿到每件商品得名称

		//属性名值列表
		gType := reflect.TypeOf(v)
		gValue := reflect.ValueOf(v)
		numField := gType.NumField()
		fmt.Println(gType.Name())
		nameValueMap := make(map[string]interface{})

		for i := 0; i < numField; i++ {
			field := gType.Field(i)
			fieldName := field.Name
			//fieldValue是一个反射Value类型，所以使用Interface方法讲Value转换成正射类型
			fieldValue := gValue.FieldByName(fieldName).Interface()
			//fmt.Println(fieldName, fieldValue)
			nameValueMap[fieldName] = fieldValue
		}
		fmt.Println()
		fmt.Println(nameValueMap)
		//写出到同名文件
		filePath := "E:\\golang\\反射\\"
		nameStr := nameValueMap["Name"].(string)
		nameStr += ".txt"
		dstFile, err := os.OpenFile(filePath+nameStr, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer func() {
			_ = dstFile.Close()
		}()
		writer := bufio.NewWriter(dstFile)
		for key, value := range nameValueMap {
			_, _ = writer.WriteString(key + fmt.Sprintf(":%v\n", value))
		}
		_ = writer.Flush()
	}
}
