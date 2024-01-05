package main

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"user-register-api/interfaces/handler"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Connection Successful",
		})
	})
	r.POST("/signin", interfaces.HandleUserSignin)

	r.Run(":8080")
}