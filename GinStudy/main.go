package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	/*
		get请求示例
	*/
	router.GET("/info", func(c *gin.Context) {
		user := c.DefaultQuery("user", "root")
		password := c.Query("password")
		c.JSON(200, gin.H{
			"code":     200,
			"message":  "success",
			"user":     user,
			"password": password,
		})
	})

	/*
		post请求示例
	*/
	router.POST("/auth", func(c *gin.Context) {
		// 获取form表单的数据
		user := c.DefaultPostForm("user", "post_user")
		password := c.PostForm("password")

		c.JSON(200, gin.H{
			"code":     200,
			"message":  "success",
			"user":     user,
			"password": password,
		})
	})

	_ = router.Run("127.0.0.1:8090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
