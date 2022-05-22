package main

import (
	"backend/handler"
	"backend/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Register_api(r)

	handler.Init_Db()
	defer handler.End_Db()

	r.Run(":8888")
}
