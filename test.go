package main

import (
	// "crypto/rand"
	"fmt"
	// "strings"
	"time"
	// "io"
)

// func main() {
// 	b := make([]byte, 8)
// 	if _, err := io.ReadFull(rand.Reader, b); err != nil {
// 		panic(err)
// 	}
// 	// rand.Read(b)
// 	a := string(b)
// 	fmt.Println(a)
// }

// type IFly interface {
// 	Fly()
// }

// type Bird struct {
// 	name string
// }

// type Plane struct {
// 	name string
// }

// func (b *Bird) Fly() {
// 	fmt.Println(b.name + " is flying...")
// }

// func (p *Plane) Fly() {
// 	fmt.Println(p.name + " is flying...")
// }
// func main() {
// 	// var fly IFly = &Bird{}   //使用IFly接口调用fly方法
// 	// var fly *Bird = new(Bird) //使用Bird对象调用fly方法
// 	var (
// 		bfly IFly = &Bird{"求是鹰"} //可以将实现接口的对象赋给该接口
// 		pfly IFly = &Plane{name: "J10B"}
// 	)
// 	bfly.Fly()
// 	pfly.Fly()
// 	bfly = pfly
// 	bfly.Fly()
// }
func main() {
	// endTime := time.Now().Unix()
	// fmt.Println(endTime)
	// timeObj := time.Unix(endTime, 0)
	// fmt.Println(timeObj)
	// endTimeStr := timeObj.Format("20060102150405")
	// time.Now().Format("20060102150405")
	p := fmt.Println
	p(time.Now().Format("20060102150405"))
}
