package main

import (
	"github.com/gin-gonic/gin"
)

type Login struct {
	USER     string `json:"user" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	r.POST("/form_post", func(c *gin.Context) {
		var loginReqPar Login
		c.BindJSON(&loginReqPar)
		c.JSON(200, gin.H{"status": "getted", "user": loginReqPar.USER})
	})
	r.Run()
}
