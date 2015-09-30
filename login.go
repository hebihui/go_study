package main

import (
	"crypto/md5"
	. "fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var table = make(map[string]int)

func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数
	Println(r.Form)
	Println("path:", r.URL.Path)
	Println("scheme", r.URL.Scheme)
	Println(r.Form["username"])
	Println(r.Form["password"])
	for k, v := range r.Form {
		Println("key:", k)
		Println("val:", strings.Join(v, ""))
	}
	Fprintf(w, "HELLO, HEBIHUI.YOU WILL SUCCESS SOONER OR LATER!")
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数
	Println("method:", r.Method)
	h := md5.New()
	test := md5.New()
	// str := "hebihui"
	// test.Write([]byte(str))
	haha := Sprintf("%x", test.Sum(nil))
	Println("my:", haha)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else {
		token := r.Form.Get("token")
		if token != "" {
			if isValidation(token) {
				table[token] = 1
				Fprintf(w, "OK.")
			} else {
				Fprintf(w, "fail.")
			}
		} else {
			Fprintf(w, "null") //不合法
		}
		Println("username:", r.Form["username"])
		Println("password:", r.Form["password"])
	}
}

func isValidation(t string) bool {
	if table[t] == 1 {
		return false
	}
	return true
}

func main() {
	http.HandleFunc("/", sayHelloWorld)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
