package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Course_Select struct {
	Student_id int    `db:"student_id" uri:"student_id" json:"student_id"`
	Course_id  string `db:"course_id" uri:"course_id" json:"course_id"`
	Teacher_id int    `db:"teacher_id" json:"teacher_id"`
	Score      int    `db:"score" json:"score"`
}

type Student_View struct {
	Student_id  int    `db:"student_id"`
	Stu_Name    string `db:"username"`
	Course_id   string `db:"course_id"`
	Teacher     string `db:"teacher"`
	Score       int    `db:"score"`
	Course_Name string `db:"name"`
}

func Select_getby_student_handler(c *gin.Context) {
	var student_info Course_Select
	var student_course []Course_Select
	var view []Student_View
	if err := c.ShouldBindJSON(&student_info); err != nil {
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

	if err := Db.Select(&view,
		`select courses.course_id, users.username, courses.teacher, courses.name, courses.score, course_select.score 
		from users, course_select, courses
		where users.id = ? 
		and course_select.student_id = users.id
		and courses.course_id = course_select.course_id
		and courses.teacher_id = course_select.teacher_id
		`, student_info.Student_id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": view})
}

func Update_Score_handler(c *gin.Context) {
	var student_info Course_Select
	var student_course []Course_Select
	if err := c.ShouldBindJSON(&student_info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if student_info.Student_id == 0 || len(student_info.Course_id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Student_id or Course_id not provided."})
		return
	}
	if err := Db.Select(&student_course,
		"select teacher_id from course_select where course_id = ? and student_id = ?",
		student_info.Course_id, student_info.Student_id); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(student_course) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Student_id = " +
			strconv.FormatInt(int64(student_info.Student_id), 10) +
			" Course_id = " + student_info.Course_id + " Not found"})
		return
	}

	if student_info.Teacher_id != student_course[0].Teacher_id {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "Auth error, error teacher_id = " +
				strconv.FormatInt(int64(student_info.Teacher_id), 10)})
		return
	}

	if student_info.Score < 0 {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "bad score = " +
				strconv.FormatInt(int64(student_info.Score), 10)})
		return
	}

	res, err := Db.Exec("update course_select set score=? where student_id = ? and course_id = ?",
		student_info.Score, student_info.Student_id, student_info.Course_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := res.RowsAffected()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": strconv.FormatInt(id, 10) + " Rows Affected\n"})
}
