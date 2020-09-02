package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

const (
	userName = "root"
	password = "123456"
	ip = "127.0.0.1"
	port = "3306"
	dbName = "test"
)
//Db数据库连接池
var DB *sql.DB

//注意方法名大写，就是public
func InitDB()  {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	//path_ := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	//如果没有密码
	path := strings.Join([]string{"root@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"},"")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil{
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")
}

func QueryUser(name string) (pass string){
	rows, err := DB.Query("select password from t1 where name = ?",name)
	if err != nil {
		fmt.Println(err.Error())
	}
	pass = "";
	for rows.Next() {
		if err := rows.Scan(&pass); err != nil {
			fmt.Println(err.Error())
		}
	}
	return
}

func InsertUser(name, pass string) (bool) {
	_,err := DB.Exec("insert into t1 values(?,?)",name,pass)
	var res bool
	if err != nil {
		res = false
		fmt.Println(err.Error())
	} else {
		res = true
	}
	return res
}