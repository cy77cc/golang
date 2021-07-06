package main

import "fmt"

func checkSortSlice(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := i+1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}

func bubbleSortSlice(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func main() {
	slice := []int{5, 2, 7, 1, 0, 9, 8, 6, 4, 3}
	checkSortSlice(slice)
	//bubbleSortSlice(slice)
	fmt.Println(slice)
}
