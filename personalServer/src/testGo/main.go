package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)


var (
	USERNAME = ""//"myuser"
	PASSWORD = ""//"mypassword"
	NETWORK = "tcp"
	SERVER = "127.0.0.1"//"49.234.102.11"
	PORT = 3306
	DATABASE = "tel"
)




func initDB() (*sql.DB) {
	if len(USERNAME) > 0 {
		USERNAME += ":"
	}
	conn := fmt.Sprintf("%s%s@%s(%s:%d)/%s?charset=utf8",USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	fmt.Println(conn)
	DB, err := sql.Open("mysql",conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return nil
	}
	DB.SetConnMaxLifetime(100*time.Second)	//最大连接周期，超时close
	DB.SetMaxOpenConns(1000)	//最大连接数
	return DB
}

func showall() {
	rows,_ :=
}

func insert(name string, age int) {

	tx,_ := DB.Begin()
	res,err :=tx.Exec("INSERT INTO tel VALUES(?,?)",name,age)
	fmt.Println(res,err)
	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
	}
}


var DB *sql.DB
func init() {
	DB = initDB()
}

func main(){
	insert("joshuasun_1",1)
	defer func() {
		DB.Close()
	}()
}
