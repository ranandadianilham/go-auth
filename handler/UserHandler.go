package handler

import (
	"fmt"
	"net/http"
	"template/go-auth/auth"
	"template/go-auth/model"

	"github.com/gin-gonic/gin"
)

func Testing(c *gin.Context) {
	c.IndentedJSON(200, "ok")
}

func SignIn(c *gin.Context) {
	c.Request.Header.Add("Content-Type", "application/json")

	var u model.User
	c.BindJSON(&u)

	fmt.Println("usr", u.Username)
	fmt.Println("pp", u.Password)

	if u.Username == "chupi" && u.Password == "123456" {
		tokenString, err := auth.CreateToken(u.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.JSON(http.StatusOK, gin.H{
			"token": tokenString,
		})
		return
	} else {
		c.JSON(http.StatusUnauthorized, "Invalid Credentials")
	}
}
