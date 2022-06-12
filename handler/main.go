package handler

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var Db *sqlx.DB

func Init_Db() {
	config := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			viper.GetString("db.username"),
			viper.GetString("db.password"),
			viper.GetString("db.addr"),
			viper.GetInt("db.port"),
			viper.GetString("db.dbname"))


	database, err := sqlx.Open("mysql", config)
	if err != nil {
		fmt.Println("open mysql failed", err)
		return
	}
	Db = database
}

func End_Db() {
	Db.Close()
}
