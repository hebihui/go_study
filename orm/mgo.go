package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

type Men struct {
	Persons []Person
}

const (
	URL = "localhost"
)

func main() {

	session, err := mgo.Dial(URL) //连接数据库
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	db := session.DB("test") //数据库名
	collection := db.C("user")

	count, err := collection.Count()
	if err != nil {
		panic(err)
	}
	fmt.Println("things object count: ", count)

	// 插入元素
	temp := &Person{
		Name:  "hebihui",
		Phone: "15088851094",
	}

	err = collection.Insert(&Person{"dawn", "88888"}, temp)
	if err != nil {
		panic(err)
	}

	//查询元素
	res := Person{}
	err = collection.Find(bson.M{"name": "dawn"}).One(&res)
	if err != nil {
		panic(err)
	}
	fmt.Println("name:", res.Name, "phone:", res.Phone)

	// _, err = collection.RemoveAll(bson.M{"name": "dawn"})
	// // if err != nil {
	// // 	panic(err)
	// // }
}
