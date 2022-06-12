package auth

import (
	"backend/handler"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type UserInfo struct {
	User_id  int    `json:"user_id"`
	Password string `json:"password"`
}

type User_Type int

const (
	Student_Type User_Type = 0
	Teacher_Type User_Type = 1
	Admin_Type   User_Type = 2
)

func (Type User_Type) String() string {
	switch Type {
	case Student_Type:
		return "Student"
	case Teacher_Type:
		return "Teacher"
	case Admin_Type:
		return "Admin"
	default:
		return "UNKONW"
	}
}

const TokenExpireDuration = time.Hour * 2

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{"error": "No token."})
			c.Abort()
			return
		}

		payload, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(os.Getenv("ACCESS_SECRET")), nil
		})

		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		claims, ok := payload.Claims.(jwt.MapClaims)
		if !(ok && payload.Valid) {
			c.JSON(http.StatusOK, gin.H{"error": "cannot convert claim to mapClaim"})
			c.Abort()
			return
		}
		c.Set("user_id", int(claims["user_id"].(float64)))
		user_id, _ := c.Get("user_id")
		log.Print("user_id=", user_id)
	}
}

func login(user_id int, password string) error {
	var user []handler.User
	if err := handler.Db.Select(&user, "select password from users where id = ?", user_id); err != nil {
		return err
	}

	if len(user) == 0 {
		return errors.New("no such user")
	}

	if password != user[0].Password {
		return errors.New("password error")
	}

	return nil
}

func CreateToken(user_id int) (string, error) {
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user_id
	atClaims["exp"] = time.Now().Add(TokenExpireDuration).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	if err != nil {
		return "", err
	}
	return token, nil
}

func AuthHandler(c *gin.Context) {
	var user UserInfo
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	if err := login(user.User_id, user.Password); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := CreateToken(user.User_id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func IsTypeMiddleWare(Type User_Type, Is bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id, ok := c.Get("user_id")
		if !ok {
			c.JSON(http.StatusOK, gin.H{"error": fmt.Errorf("no user_id in middleware").Error()})
			c.Abort()
			return
		}

		var user []handler.User
		if err := handler.Db.Select(&user, "select type from users where id = ?", user_id); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if len(user) == 0 {
			c.JSON(http.StatusOK, gin.H{"error": fmt.Errorf("no such user:%d", user_id).Error()})
			c.Abort()
			return
		}

		result := (user[0].Type != int(Type))
		if !Is {
			result = !result
		}

		if result {
			c.JSON(http.StatusOK, gin.H{"error": fmt.Errorf("except type:%v, user type:%v", Type, user[0].Type).Error()})
			c.Abort()
			return
		}
	}
}
