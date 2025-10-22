package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func useMidware() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	router.Run(":8081")
}

func main() {
	//useMidware()
	//return

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

	// 输出格式
	//router.GET("test-output/:type", func(c *gin.Context) {
	//	_type := c.Param("type")
	//	switch _type {
	//	case "1":
	//		c.JSON(200, gin.H{
	//			"message": "It's JSON",
	//		})
	//	case "2":
	//		c.XML(200, gin.H{
	//			"message": "It's XML",
	//		})
	//	case "3":
	//		c.YAML(200, gin.H{
	//			"message": "It's YAML",
	//			"code":    200,
	//		})
	//	}
	//})

	//router.LoadHTMLGlob("templates/*")
	//router.GET("/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	//		"title": "Home Page",
	//	})
	//})

	//router.LoadHTMLGlob("templates/**/*")
	//router.GET("/post/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "post/index.tmpl", gin.H{
	//		"title": "Post Index",
	//	})
	//})

	// 参数绑定
	type formA struct {
		Name string `json:"name" xml:"name" form:"name" binding:"required"`
		Foo  string `json:"foo" xml:"foo" form:"foo" binding:"required"`
	}
	type formB struct {
		Name string `json:"name" xml:"name" form:"name" binding:"required"`
		Bar  string `json:"bar" xml:"bar" form:"bar" binding:"required"`
	}
	router.POST("/login/:foo", func(c *gin.Context) {
		//ObjA := formA{}
		//ObjB := formB{}
		//if errA := c.ShouldBind(&ObjA); errA == nil {
		//	c.String(http.StatusOK, "The body should be formA")
		//} else if errB := c.ShouldBind(&ObjB); errB == nil {
		//	c.String(http.StatusOK, "The body should be formB")
		//} else {
		//	fmt.Println("objA, objB:", ObjA, ObjB)
		//	fmt.Println("errA, errB:", errA, errB)
		//}

		//if errA := c.ShouldBindBodyWith(&ObjA, binding.JSON); errA == nil {
		//	c.String(http.StatusOK, "The body should be formA")
		//} else if errB := c.ShouldBindBodyWith(&ObjB, binding.JSON); errB == nil {
		//	c.String(http.StatusOK, "The body should be formB")
		//} else {
		//	fmt.Println("objA, objB:", ObjA, ObjB)
		//	fmt.Println("errA, errB:", errA, errB)
		//}

		//c.MustBindWith(&ObjA, binding.JSON) // 如果传的数据中缺少formA必填的则退出请求，返回400状态码

		// 非绑定参数
		//f1 := c.Query("foo")
		//f2 := c.Param("foo")
		//f3 := c.PostForm("foo")
		//fmt.Println(f1, f2, f3)
		//c.JSON(200, gin.H{
		//	"f1": f1,
		//	"f2": f2,
		//	"f3": f3,
		//})
	})

	// 中间件
	router.GET("/test-midware", mw1(), mw2(), func(c *gin.Context) {
		fmt.Println("self")
		c.String(200, "self")
	})
	secres := gin.H{
		"foo":      gin.H{"name": "foo", "email": "foo@qq.com", "phone": "14547878787"},
		"xiaoming": gin.H{"name": "xiaoming", "email": "manu@qq.com", "phone": "14547878787"},
		"manu":     gin.H{"name": "", "email": "@qq.com", "phone": "14547878787"},
	}
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":      "bar",
		"xiaoming": "123456",
		"manu":     "iloveyou",
		"lemon":    "888",
	}))
	authorized.GET("/secrets", mw1(), func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		if secrets, ok := secres[user]; ok {
			c.JSON(200, gin.H{"user": user, "secrets": secrets})
		} else {
			c.JSON(200, gin.H{"user": user, "secrets": "No secrets"})
		}
	})

	//验证器
	type Booking struct {
		CheckIn  time.Time `json:"check_in" xml:"check_in" form:"check_in" binding:"required" time_format:"2006-01-02"`
		CheckOut time.Time `json:"check_out" xml:"check_out" form:"check_out" binding:"required,gtfield=CheckIn,bookabledate" time_format:"2006-01-02" `
	}

	var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
		data, ok := fl.Field().Interface().(time.Time)
		if ok {
			today := time.Now()
			if today.After(data) {
				return false
			}
		}
		return true
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}
	router.GET("/bookings", func(c *gin.Context) {
		var booking Booking
		if err := c.ShouldBindWith(&booking, binding.Query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": booking, "message": "It's valid"})
		}
	})
	router.Run(":8080")
}

func mw1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before mw1")
		c.Next()
		fmt.Println("After mw1")
	}
}

func mw2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before mw2")
		c.Next()
		fmt.Println("After mw2")
	}
}
