package main

import (
	"crypto/md5"
	"database/sql"
	. "fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type myMux struct{} //自定义路由

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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func databaseOperation(w http.ResponseWriter, r *http.Request) {
	Println("method:", r.Method)
	if r.Method == "GET" {
		db, err := sql.Open("mysql", "root:@/test?charset=utf8")
		checkErr(err)

		//插入数据
		// stmt, err := db.Prepare("insert userinfo set username=?,department=?,created=?")
		// checkErr(err)

		// res, err := stmt.Exec("hebihui", "98510特混旅", "2016-07-01")
		// checkErr(err)
		//直接插入数据
		// res, err := db.Exec("insert userinfo set username=?,department=?,created=?", "hebihui", "98510特混旅", "2016-07-01")
		// checkErr(err)

		// id, err := res.LastInsertId()
		// checkErr(err)

		// Println(id)
		// Fprintf(w, "%d", id)

		//更新数据
		// stmt1, err1 := db.Prepare("update userinfo set username=? where uid=?")
		// checkErr(err)

		// res1, err1 := stmt1.Exec("hebihui", id)
		// checkErr(err)

		// affect1, err1 := res1.RowsAffected()
		// checkErr(err1)

		// Println(affect1)

		//查询数据
		rows, err := db.Query("select * from userinfo")
		checkErr(err)

		for rows.Next() {
			var uid int
			var username string
			var department string
			var created string
			err = rows.Scan(&uid, &username, &department, &created)
			checkErr(err)
			Println(uid)
			Println(username)
			Println(department)
			Println(created)
		}

		//删除数据
		// stmt3, err3 := db.Prepare("delete from userinfo where uid=?")
		// checkErr(err3)

		// res3, err3 := stmt3.Exec(id)
		// checkErr(err3)

		// affect3, err3 := res3.RowsAffected()
		// checkErr(err3)

		// Println(affect3)

		db.Close()
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	Println("method:", r.Method)
	r.ParseForm()
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else if r.Method == "POST" {
		// 解析账号密码
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		// Println(username)

		//查询数据库
		db, err := sql.Open("mysql", "root:@/test?charset=utf8")
		checkErr(err)

		var sql string = "select username,password,created from userandpassword where username=" + "'" + username + "'"
		Println(sql)
		rows, err := db.Query(sql)
		checkErr(err)

		var d_username string
		var d_password string
		var created string
		for rows.Next() {
			err = rows.Scan(&d_username, &d_password, &created)
			checkErr(err)
			Println(d_password)
			Println(created)
		}

		if password == d_password {
			Fprintf(w, "%s", "登陆成功。")
		} else {
			Fprintf(w, "%s", "账号或者密码错误！")
		}
	} else {
		Fprintf(w, "%s", "非法请求！")
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := Sprintf("x", h.Sum(nil))

		t, _ := template.ParseFiles("front.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			Println(err)
			return
		}
		defer file.Close()
		Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	Fprintf(w, "自定义路由。")
}

func (m *myMux) ServeHTTP(w http.ResponseWriter, r *http.Request) { //自定义路由实现Hander接口的ServeHTTP方法
	if r.URL.Path == "/" {
		sayHello(w, r) //调用自定义路由对应的处理函数
		return
	}
	if r.URL.Path == "/upload" {
		upload(w, r)
		return
	}
	if r.URL.Path == "/databaseOperation" {
		databaseOperation(w, r)
		return
	}
	if r.URL.Path == "/login" {
		login(w, r)
		return
	}
	http.NotFound(w, r) //匹配失败，调用http.NotFound
	return
}

func main() {
	// http.HandleFunc("/", sayHelloWorld)      //默认路由
	mux := &myMux{}
	err := http.ListenAndServe(":9999", mux) //设置监听端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
