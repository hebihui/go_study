package main

import (
	"fmt"
)

func insertSort(buf []int) []int {
	len := len(buf)
	for i := 1; i < len; i++ {
		tmp := buf[i]
		j := i - 1
		for ; j >= 0; j-- {
			if buf[j] > tmp {
				buf[j+1] = buf[j]
			} else {
				buf[j+1] = tmp
				break
			}
		}
		buf[j+1] = tmp
	}
	return buf
}

func main() {
	arr := []int{4, 3, 1, 341, 31, 41, 14}
	res := insertSort(arr)
	fmt.Println(res)
}
