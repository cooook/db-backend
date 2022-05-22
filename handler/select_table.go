package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Course_Select struct {
	Student_id 	int    `db:"student_id" uri:"student_id" json:"student_id"`
	Course_id  	string `db:"course_id" uri:"course_id" json:"course_id"`
	Teacher_id  string `db:"teacher_id" json:"teacher_id"`
	Score     	int    `db:"score" json:"score"`
}

type Student_View struct {
	Student_id 	int    	`db:"student_id"`
	Stu_Name	string 	`db:"user_name"`
	Course_id 	string 	`db:"course_id"`
	Teacher 	string 	`db:"teacher"`
	Score 		int		`db:"score"`
	Course_Name string 	`db:"name"`
}

func Selcet_getby_student_handler(c *gin.Context) {
	var student_info Course_Select
	var student_course []Course_Select
	if err := c.ShouldBindUri(&student_info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := Db.Select(&student_course, "select * from course_select where student_id=?", student_info.Student_id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(student_course) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No student with student_id=" + strconv.FormatInt(int64(student_info.Student_id), 10)})
		return
	}
	if err := Db.Select(&)
}
