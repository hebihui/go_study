package main

import (
	"fmt"
)

func hano(x int) int64 {
	if x == 1 {
		return 1
	} else {
		return 2*hano(x-1) + 1
	}
}

func main() {
	// var num int = 5
	num1 := 10
	fmt.Println(hano(num1))
}
