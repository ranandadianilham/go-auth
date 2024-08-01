package main

import (
	"template/go-auth/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/login", handler.SignIn)

	router.Run("localhost:8080")
}
