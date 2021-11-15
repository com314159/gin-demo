package main

import "fmt"
import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()
	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.GET("first", func(c *gin.Context) {
		fmt.Println("first .........")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
