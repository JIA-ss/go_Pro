package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db,err := sql.Open("mysql","")
	if err != nil {
		fmt.Println("open db err: ",err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("ping db err",err)
	}



}