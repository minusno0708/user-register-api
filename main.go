package main

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"user-register-api/model"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Connection Successful",
		})
	})
	r.POST("/signin", func(c *gin.Context) {
		var user model.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Body does not exist",
			})
			return
		}
		if user.UserID == "" || user.Password == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Body is not valid",
			})
			return
		}
		if user.Username == "" {
			user.Username = user.UserID
		}

		db, err := model.ConnectDB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Database connection error",
			})
			return
		}
		defer db.Close()

		err = model.InsertUser(db, user)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"message": "User already exists",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "User created successfully",
			"user": user,
		})
	})

	r.Run(":8080")
}