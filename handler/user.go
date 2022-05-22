package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	UserId     int    `db:"id" json:"user_id" uri:"user_id" binding:"required"`
	Username   string `db:"username" json:"username"`
	Password   string `db:"password" json:"password"`
	College    string `db:"college" json:"college"`
	Profession string `db:"profession" json:"profession"`
	Type       int    `db:"type" json:"type"`
}

func User_get_all(c *gin.Context) {
	var user []User
	if err := Db.Select(&user, "select id, username, college, profession from users"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(user) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Such student"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func User_get_handler(c *gin.Context) {
	var user []User
	var json User
	if err := c.ShouldBindUri(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := Db.Select(&user, "select id, username, college, profession from users where id=?", json.UserId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(user) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Such student"})
		return
	}
	c.JSON(http.StatusOK, user[0])
}

func User_post_handler(c *gin.Context) {
	var json User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() + " Parse JSON"})
		return
	}

	r, err := Db.Exec("insert into users(id, username, password, college, profession, type)values(?, ?, ?, ?, ?, ?)",
		json.UserId, json.Username, json.Password, json.College, json.Profession, json.Type)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = r.LastInsertId()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "Add user success!"})
}

func User_put_handler(c *gin.Context) {
	var json User
	var old []User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't parse json"})
		return
	}

	if err := Db.Select(&old, "select college, profession, password from users where id=?", json.UserId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(old) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Such student"})
		return
	}

	if json.College == "" {
		json.College = old[0].College
	}
	if json.Profession == "" {
		json.Profession = old[0].Profession
	}
	if json.Password == "" {
		json.Password = old[0].Password
	}

	res, err := Db.Exec("update users set college=?, profession=?, password=? where id=?",
		json.College, json.Profession, json.Password, json.UserId)
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
