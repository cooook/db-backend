package router

import (
	"backend/auth"
	"backend/handler"

	"github.com/gin-gonic/gin"
)

func Register_api(r *gin.Engine) {
	r.POST("/auth", auth.AuthHandler)

	api_group := r.Group("/v1").Use(auth.JWTAuth())

	api_group.GET("/users", auth.IsTypeMiddleWare(auth.Admin_Type, true), handler.User_get_all)           //.Use(auth.IsTypeMiddleWare(auth.Student_Type, true))
	api_group.POST("/users", auth.IsTypeMiddleWare(auth.Admin_Type, true), handler.User_post_handler)     // .Use(auth.IsTypeMiddleWare(auth.Admin_Type, true))
	api_group.PUT("/users", auth.IsTypeMiddleWare(auth.Admin_Type, true), handler.User_put_handler)       //.Use(auth.IsTypeMiddleWare(auth.Admin_Type, true))
	api_group.POST("/courses", auth.IsTypeMiddleWare(auth.Admin_Type, true), handler.Course_post_handler) //.Use(auth.IsTypeMiddleWare(auth.Admin_Type, true))

	api_group.GET("/users/:user_id", handler.User_get_handler)
	api_group.GET("/courses", handler.Course_get_all)
	api_group.GET("/courses/:course_id", handler.Course_get_handler)

	api_group.GET("/course_table", handler.Select_getby_student_handler)
	api_group.PUT("/course_table", handler.Update_Score_handler)
	api_group.POST("/course_table", handler.Add_course_select_entry)
}
