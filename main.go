package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Connection Successful",
		})
	})
	r.POST("/signin", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad Request",
			})
			return
		}
		if user.UserID == "" || user.Password == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		if user.Username == "" {
			user.Username = user.UserID
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "User created successfully",
			"user": user,
		})
	})

	r.Run(":8080")
}