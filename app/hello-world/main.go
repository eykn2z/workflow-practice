package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Fatalln("エラーで落とす")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
	r.Run(":80")
}
