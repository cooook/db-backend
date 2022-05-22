package handler

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func Init_Db() {
	database, err := sqlx.Open("mysql", "backend:123456@tcp(localhost:3306)/course_select")
	if err != nil {
		fmt.Println("open mysql failed", err)
		return
	}
	Db = database
}

func End_Db() {
	Db.Close()
}
