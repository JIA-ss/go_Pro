package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("admin")
	password := r.FormValue("password")
	//fmt.Println("username: " + user)
	//fmt.Println("password: "+password)
	p := QueryUser(user)
	if p == "" {
		w.Write([]byte("用户名不存在！"))
		return
	}
	if p != password {
		w.Write([]byte("密码错误"))
		return
	} else {
		w.Write([]byte("登录成功"))
		return
	}
}

func Signup(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("admin")
	password := r.FormValue("password")
	//fmt.Println("username: " + user)
	//fmt.Println("password: "+password)
	if user == "" || password == "" {
		w.Write([]byte("用户名 和 密码不能为空"))
		return
	}
	p := QueryUser(user)
	if p != "" {
		w.Write([]byte("用户名已被占用"))
		return
	}
	if InsertUser(user,password) {
		w.Write([]byte("sign up succeed! user: " + user + "\tpassword: " + password))
	} else {
		w.Write([]byte("sign up failed! user: " + user + "\tpassword: " + password))
	}
	return
}

func main()  {
	InitDB()
	defer func(){
		DB.Close()
	}()

	http.HandleFunc("/",func(res http.ResponseWriter, req *http.Request) {
		t, err := template.ParseFiles("./templates/login.html")
		if err != nil {
			log.Println(err.Error())
		}
		t.Execute(res,nil)
	})
	http.HandleFunc("/Login",Login)
	http.HandleFunc("/Signup",Signup)
	fmt.Println("start http server")
	http.ListenAndServe(":8080",nil)
}
