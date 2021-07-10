package main

import "fmt"

func main() {
	data := make([]interface{}, 0)
	data = append(data, "小爱")
	data = append(data, 18)
	data = append(data, 1.2)
	data = append(data, false)
	data = append(data, '0')
	for i := 0; i < len(data); i++ {
		switch data[i].(type) {
		case int:
			fmt.Println("发现int类型")
		case string:
			fmt.Println("发现字符串类型")
		case float64:
			fmt.Println("发现float64类型")
		case bool:
			fmt.Println("发现bool类型")
		default:
			fmt.Println("发现其他类型")
		}
	}
	for _, v := range data {
		if s, ok := v.(string); ok {
			fmt.Println("找到了string类型", s)
		} else {
			fmt.Println("没有找到")
		}

	}
}
