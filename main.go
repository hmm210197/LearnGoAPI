package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":        "posted",
			"messageGetted": message,
			"nickGetted":    nick,
		})
	})
	r.Run()
}
