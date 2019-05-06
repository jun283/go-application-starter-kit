package main

import (
	"database/sql"
	"fmt"
)

var (
	//live DB connection string
	DBConn_live = ""

	//debug DB connection string
	DBConn_debug = ""
)

func DBConnect() (db *sql.DB, err error) {

	//connect to mysql server
	var DBConn string

	if conf.Debug {
		DBConn = DBConn_debug
	} else {
		DBConn = DBConn_live
	}

	db, err = sql.Open("mysql", DBConn)
	if err != nil {
		errLog.Println(err)
		return
	}
	//defer db.Close()

	fmt.Println("Database connect successful")

	return
}
