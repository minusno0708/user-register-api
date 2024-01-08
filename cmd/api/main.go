package main

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"user-register-api/interfaces/handler"
	"user-register-api/usecase"
	"user-register-api/infrastructure/persistence"
)

func main() {
	userPersistence := persistence.NewUserPersistence()
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Connection Successful",
		})
	})
	r.POST("/signin", userHandler.HandleUserSignin)

	r.Run(":8080")
}