package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	path2 "path"
	"strings"
	"testProject/20200926-mysql/dbutil"
	"time"
)

func RenderTemplate(path string) string {
	realPath := path2.Join("./templates", path)
	content, err := ioutil.ReadFile(realPath)
	if err != nil {
		fmt.Println("read file failed, err:", err)
		panic(err)
	}
	return string(content)
}

type UserInfo struct {
	id       int
	username string
	password string
}
type MyHandler struct{}

func doGet(w *http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/regist":
		fmt.Fprint(*w, RenderTemplate("regist.html"))
	case "/login":
		fmt.Fprint(*w, RenderTemplate("login.html"))
	}
}
func regist(username string, password string) (msg string) {
	msg = "注册失败"
	sqlStr := "insert into userinfo(username,password)values(?,?)"
	stmt, err := dbutil.DB.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(username, password)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	_, err = result.LastInsertId()
	if err == nil {
		msg = "注册成功"
		return
	}
	return
}
func login(username string, password string) (msg string) {
	msg = "登录失败"
	sqlStr := "select id,username,password from userinfo where username=? and password=?"
	stmt, err := dbutil.DB.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(username, password)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	if rows.Next() {
		msg = "登陆成功"
		return
	} else {
		return
	}
}
func doPost(w *http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	var respContent string
	switch r.URL.Path {
	case "/regist":
		respContent = regist(username, password)
	case "/login":
		respContent = login(username, password)
	}
	fmt.Fprint(*w, respContent)
}

func (myHandler *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reqMethod := strings.ToUpper(r.Method)
	switch reqMethod {
	case "GET":
		doGet(&w, r)
	case "POST":
		doPost(&w, r)
	}
}

func main() {
	// 注意，由于读取文件使用的相对路径是相对当前文件的，所以执行此程序时需要切换工作目录到此文件所在目录
	myHandler := new(MyHandler)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
