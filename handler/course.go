package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Course struct {
	Course_id  string `db:"course_id" json:"course_id" uri:"course_id" binding:"required"`
	Teacher_id int    `db:"teacher_id" json:"teacher_id" uri:"teacher_id"`
	Teacher    string `db:"teacher" json:"teacher"`
	Name       string `db:"name" json:"name"`
	Max_num    int    `db:"max_num" json:"max_num"`
	Num        int    `db:"num" json:"num"`
	Score      int    `db:"score" json:"score"` // 学分
}

func Course_get_all(c *gin.Context) {
	var courses []Course

	if err := Db.Select(&courses, "select name, course_id, teacher, num, max_num from courses"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func Course_get_handler(c *gin.Context) {
	var json Course
	var courses []Course
	if err := c.ShouldBindUri(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := Db.Select(&courses, "select name, course_id, teacher, num, max_num from courses where course_id=?",
		json.Course_id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, courses)
}

func Course_post_handler(c *gin.Context) {
	var json Course
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r, err := Db.Exec("insert into courses(course_id, teacher_id, teacher, name, max_num, num, score)values(?, ?, ?, ?, ?, ?, ?)",
		json.Course_id, json.Teacher_id, json.Teacher, json.Name, json.Max_num, 0, json.Score)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	line, err := r.LastInsertId()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "Insert after " + strconv.FormatInt(line, 10)})
}
