package main

import (
	"github.com/eric-chao/go-micro/gin-demo/handler"
	model "github.com/eric-chao/go-micro/gin-demo/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.StaticFile("/tmp/test.json", "./tmp/test.json")
	r.Static("/static", "/Users/zhxj/Downloads")
	r.StaticFS("/static1", gin.Dir("/Users/zhxj/Downloads", true))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Home Page")
	})

	adminGroup := r.Group("/admin")
	adminGroup.Use(gin.BasicAuth(gin.Accounts{"admin": "admin"}))
	adminGroup.GET("/index", func(c *gin.Context){
		c.JSON(200, "Admin Page")
	})

	// auth
	r.Use(gin.BasicAuth(gin.Accounts{"Admin": "123456"}))

	r.GET("/protobuf", func(c *gin.Context) {
		data := &model.User{
			Name: "张三",
			Age:  20,
		}
		c.ProtoBuf(http.StatusOK, data)
	})

	r.GET("/format", func(c *gin.Context) {
		data := &model.User{
			Name: "张三",
			Age:  20,
		}

		if c.ContentType() == "application/json" {
			c.JSON(http.StatusOK, data)
		} else {
			c.ProtoBuf(http.StatusOK, data)
		}
	})

	r.GET("/binding", handler.BindingTest)
	r.Run(":8080")
}
