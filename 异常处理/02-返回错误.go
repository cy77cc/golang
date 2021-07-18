package main

import (
	"fmt"
	"math"
)

type userDefinedError struct {}

func (u *userDefinedError) Error() string {
	return "半径不能小于0"
}

func getCircleArea(radius float64) (float64, error) {
	if radius < 0 {
		//return 0, errors.New("半径不能为负数")
		return 0, new(userDefinedError)
	}

	return math.Pi * radius * radius, nil
}

func main() {
	//fmt.Println(getCircleArea(-3))
	area, err := getCircleArea(-2)
	fmt.Println(area)
	fmt.Println(err)
}
