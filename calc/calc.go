package main

import (
	"add"
	"fmt"
	"os"
	"sqrt"
	"strconv"
)

var Usage = func() { //把匿名函数作为一个变量使用
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\nThe commmands are:\n\tAddition of two values.\n\tsqrt\tSquare root of non-negative value.")
}

func main() {
	args := os.Args                   //使用os包读取输入参数
	if args == nil || len(args) < 2 { //若输入参数不符合要求，则提示用法
		Usage()
		return
	}
	switch args[1] { //根据第一个参数判断执行什么操作
	case "add":
		if len(args) != 4 {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}
		v1, err1 := strconv.Atoi(args[2]) //使用sreconv提供的atoi方法把字符型参数转化为数字
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}
		ret := add.Add(v1, v2)
		fmt.Println("res:", ret)
	case "sqrt":
		if len(args) != 3 {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		v, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		ret := sqrt.Sqrt(v)
		fmt.Println("res:", ret)
	default:
		Usage()
	}
}
