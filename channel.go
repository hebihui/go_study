package main

import (
	"fmt"
	"time"
)

// var count int = 0

// func Count(ch chan int) {
// 	println("count")
// 	ch <- 1
// }

// func main() {
// 	chs := make([]chan int, 10)
// 	for i := 0; i < 10; i++ {
// 		chs[i] = make(chan int)
// 		fmt.Println(i)
// 		go Count(chs[i])
// 	}
// 	for _, ch := range chs {
// 		t := <-ch
// 		count += t
// 		fmt.Println(count)
// 	}
// }

/*使用timeout防止死锁*/
func Count(ch chan int, v int) {
	ch <- v
	fmt.Println("count")
}
func main() {
	ch := make(chan int, 1024)
	for i := 0; i < 10; i++ {
		go Count(ch, i)
	}
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-timeout:
			goto label
		}
	}
label:
	fmt.Println("end")
}
