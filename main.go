package main

import (
	"backend/config"
	"backend/handler"
	"backend/router"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var cfg = pflag.StringP("config", "c", "", "apiserver config file path.")

func main() {

	pflag.Parse()
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	log.Printf("The config inited.")

	handler.Init_Db()
	defer handler.End_Db()

	r := gin.Default()
	router.Register_api(r)

	r.Run(viper.GetString("addr"))
}

// [TODO] 教师为每位同学修改成绩
