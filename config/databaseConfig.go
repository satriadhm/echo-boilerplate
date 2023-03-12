package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func DatabaseConfig() error {
	Dbconn, err := sql.Open("mysql", "root:satria2133@tcp(127.0.0.1:3306)/test1")
	if err != nil {
		fmt.Println(err.Error())
		return err
	} else {
		fmt.Println("db is connected")
	}
	// make sure connection is available
	err = Dbconn.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	Db = Dbconn
	return nil
}
