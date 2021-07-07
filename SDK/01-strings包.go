package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "html"
	fmt.Println(strings.Contains(str, "h"))
	fmt.Println(strings.ContainsRune(str, 'h'))
	fmt.Println("Hello, World")
	fmt.Printf("%c", 123)
}
