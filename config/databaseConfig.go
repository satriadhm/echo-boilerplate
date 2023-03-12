package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func DatabaseConfig() {
	Db, err := sql.Open("mysql", "root:satria2133@tcp(127.0.0.1:3306)/test1")
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("db is connected")
	}
	// make sure connection is available
	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print(Db)

}
