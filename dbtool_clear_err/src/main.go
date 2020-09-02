package main

import (
	"database/sql"
	"errors"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"time"
)


var (
	USERNAME = "myuser"
	PASSWORD = "mypassword"
	NETWORK = "tcp"
	SERVER = "9.134.233.66"
	PORT = 3306
	DATABASE = "ckv_plus_portal"
)

type FIX_INFO struct {
	Id 			int
	Bid 		string
	Region 		string
	Set 		string
	Bs1Name		string
	Bs2Name		string
	State      	int
	Num         int
}

func err_condition(fix_info FIX_INFO) bool {
	if len(fix_info.Region) != 0 && len(fix_info.Set) != 0 && len(fix_info.Bid) != 0 && len(fix_info.Bs1Name) == 0 {
		return true
	}
	return false
}

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

func get_state_err(DB *sql.DB) []FIX_INFO {
	rows, err := DB.Query("select id,bid,region,`set`,bs1_name,bs2_name from bid_apply_info where state = -3")
	if err == nil {
		errors.New("query incur error")
	}

	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	var fix_info_list  []FIX_INFO
	var fix_info FIX_INFO
	for rows.Next() {
		err = rows.Scan(&fix_info.Id, &fix_info.Bid, &fix_info.Region, &fix_info.Set, &fix_info.Bs1Name, &fix_info.Bs2Name)  //不scan会导致连接不释放
		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
		}
		if err_condition(fix_info) {
			fix_info_list = append(fix_info_list,fix_info)
		}
	}
	return fix_info_list
}

func countNum(DB *sql.DB) {
	tx,_ := DB.Begin()
	for i,_ := range fix_info_list_get{
		judge_exist_state3 := fmt.Sprintf("SELECT COUNT(*) FROM bid_apply_info where bid='%s' and state=3 and region = '' and `set` = ''",fix_info_list_get[i].Bid)
		//fmt.Println(judge_exist_state3)
		err := tx.QueryRow(judge_exist_state3).Scan(&(fix_info_list_get[i].Num))
		//fmt.Println("bid:",fix_info_list_get[i].Bid,"num:",fix_info_list_get[i].Num)
		if err != nil {
			fmt.Println(err)
		}
	}
	tx.Commit()
}


func opt_info(DB *sql.DB, fix_info FIX_INFO) {

	tx,_ := DB.Begin()
	tx.Exec("DELETE FROM bid_apply_info WHERE `set` = ? and region = ? and bid = ? and state = -3",fix_info.Set,fix_info.Region,fix_info.Bid)
	tx.Exec("UPDATE bid_apply_info SET `set` = ?,region = ? WHERE bid = ? and state = 3 and `set` = '' and region = ''",fix_info.Set,fix_info.Region,fix_info.Bid)
	tx.Commit()
}

func showByBid(DB *sql.DB, err_item FIX_INFO) {
	showStr := fmt.Sprintf("select id,bid,region,`set`,bs1_name,bs2_name,state from bid_apply_info where bid = '%s'",err_item.Bid)
	rows, err := DB.Query(showStr)
	if err == nil {
		errors.New("query incur error")
	}

	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	fmt.Println("=== current err item -> bid: ",err_item.Bid,"\tregion: ",err_item.Region,"\tset: ",err_item.Set , " ===")
	i := 1
	var fix_info FIX_INFO
	for rows.Next() {
		err = rows.Scan(&fix_info.Id, &fix_info.Bid, &fix_info.Region, &fix_info.Set, &fix_info.Bs1Name, &fix_info.Bs2Name, &fix_info.State)  //不scan会导致连接不释放

		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
		}
		fmt.Printf("-> state: %d \t region= %v \t set= %v \t bs1_name= %v \t bs2_name=%v \n",fix_info.State,fix_info.Region,fix_info.Set,fix_info.Bs1Name,fix_info.Bs2Name)
		i++
	}
}



var fix_info_list_get []FIX_INFO
func main(){

	DB := initDB()
	defer func() {
		DB.Close()
	}()
	fix_info_list_get = get_state_err(DB)
	countNum(DB)
	for _,i := range fix_info_list_get {
		if i.Num > 0 {
			showByBid(DB,i)
			var doOpt string
			fmt.Println("weather to merge and delete the error one ? (y:yes or n:no)")
			fmt.Scanln(&doOpt)
			if doOpt == "y" {
				opt_info(DB,i)
			}
			//opt_info(DB,i)
		}
	}




}
