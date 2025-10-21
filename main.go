package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// RESTFUL
	// 实际应用中uri为同一个才符合restful api
	//router.GET("/get", func(c *gin.Context) {
	//	c.String(200, "Hello World (GET)")
	//})
	//router.POST("/post", func(c *gin.Context) {
	//	c.String(200, "Hello World (POST)")
	//})
	//router.PUT("/put", func(c *gin.Context) {
	//	c.String(200, "Hello World (PUT)")
	//})
	//router.DELETE("/delete", func(c *gin.Context) {
	//	c.String(200, "Hello World (DELETE)")
	//})
	//router.HEAD("/head", func(c *gin.Context) {
	//	c.String(200, "Hello World (HEAD)")
	//})
	//router.PATCH("/patch", func(c *gin.Context) {
	//	c.String(200, "Hello World (PATCH)")
	//})

	// 路由分组
	//{
	//	v1 := router.Group("/v1")
	//	v1.GET("/login", func(c *gin.Context) {
	//		c.JSON(200, gin.H{
	//			"message": "v1 login",
	//		})
	//	})
	//}
	//{
	//	v2 := router.Group("/v2")
	//	v2.GET("/login", func(c *gin.Context) {
	//		c.JSON(200, gin.H{
	//			"message": "v2 login",
	//		})
	//	})
	//}

	// 重定向
	//router.GET("/redirect1", func(c *gin.Context) {
	//	c.Redirect(http.StatusPermanentRedirect, "http://www.baidu.com")
	//})
	//
	//router.GET("/redirect-distance", func(c *gin.Context) {
	//	c.String(200, "Hello REDIRECT")
	//})
	//router.GET("/redirect2", func(c *gin.Context) {
	//	c.Redirect(http.StatusFound, "/redirect-distance")
	//})
	//router.GET("/redirect3", func(c *gin.Context) {
	//	c.Request.URL.Path = "/redirect-distance"
	//	router.HandleContext(c)
	//})

	//静态文件
	//router.Static("/static", "./static")
	//router.StaticFile("/icon.foo", "./static/temperature.svg")
	//router.StaticFS("/static", http.Dir("static")) // "./static/sub-static”中的静态文件也能访问了

	router.GET("test-output/:type", func(c *gin.Context) {
		_type := c.Param("type")
		switch _type {
		case "1":
			c.JSON(200, gin.H{
				"message": "It's JSON",
			})
		case "2":
			c.XML(200, gin.H{
				"message": "It's XML",
			})
		case "3":
			c.YAML(200, gin.H{
				"message": "It's YAML",
				"code":    200,
			})
		}
	})

	router.Run(":8080")
}
