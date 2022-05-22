package router

import (
	"backend/handler"

	"github.com/gin-gonic/gin"
)

func Register_api(r *gin.Engine) {
	api_group := r.Group("/api")

	api_group.GET("/users", handler.User_get_all)
	api_group.GET("/users/:user_id", handler.User_get_handler)
	api_group.POST("/users", handler.User_post_handler)
	api_group.PUT("/users", handler.User_put_handler)

	api_group.GET("/courses", handler.Course_get_all)
	api_group.GET("/courses/:course_id", handler.Course_get_handler)
	api_group.POST("/courses", handler.Course_post_handler)
}
