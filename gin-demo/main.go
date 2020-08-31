package main

import (
	model "github.com/eric-chao/go-micro/gin-demo/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/protobuf", func(c *gin.Context) {
		data := &model.User{
			Name: "张三",
			Age:  20,
		}
		c.ProtoBuf(http.StatusOK, data)
	})

	r.Run(":8080")
}
