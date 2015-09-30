package main

import (
	"fmt"
)

var arr [20][20]int
var len1 int
var len2 int

func lsc(str1, str2 string) [20][20]int {
	len1 = len(str1)
	len2 = len(str2)
	max := Max(len1, len2)
	for i := 0; i < max; i++ {
		arr[0][i] = 0
		arr[i][0] = 0
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if str1[i-1] == str2[j-1] {
				arr[i][j] = arr[i-1][j-1] + 1
			} else {
				arr[i][j] = Max(arr[i-1][j], arr[i][j-1])
			}
		}
	}
	return arr
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	var res int = 0
	str1 := "cdhoan" //匹配串
	str2 := "cdhoan" //模式串
	t := lsc(str1, str2)
	for i := 0; i <= len1; i++ {
		for j := 0; j <= len2; j++ {
			if t[i][j] > res {
				res = t[i][j]
			}
		}
	}
	fmt.Println(res)
}
