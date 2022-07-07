package main

import "fmt"


func printSlice[T any](s []T) {
	for _, v := range s {
		fmt.Printf("%v\n", v)
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	printSlice[int](s)
}
