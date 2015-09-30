package main

import(
	. "fmt"
	  "net/http"
	  "strings"
	  "log"
)

type myMux struct{}                          //自定义路由

// func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()                            //解析参数
// 	Println(r.Form)
// 	Println("path:", r.URL.Path)
// 	Println("scheme", r.URL.Scheme)
// 	Println(r.Form["id"])
// 	Println(r.Form["name"])
// 	for k, v := range r.Form {
// 		Println("key:", k)
// 		Println("val:", strings.Join(v, ""))
// 	}
// 	Fprintf(w, "HELLO, HEBIHUI.YOU WILL SUCCESS SOONER OR LATER!")
// }

func sayHello(w http.ResponseWriter, r *http.Request){
	Fprintf(w, "自定义路由。")
}

func (m *myMux) ServeHTTP(w http.ResponseWriter, r *http.Request){    //自定义路由实现Hander接口的ServeHTTP方法
	if r.URL.Path == "/" {
		sayHello(w , r)                                               //调用自定义路由对应的处理函数
		return
	}
	http.NotFound(w, r)                                               //匹配失败，调用http.NotFound
	return
}

func main() {
	// http.HandleFunc("/", sayHelloWorld)      //默认路由
	mux := &myMux{}
	err := http.ListenAndServe(":9999", mux) //设置监听端口
	if err != nil{
		log.Fatal("ListenAndServe: ", err)
	}
}
