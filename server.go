package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-crud-mysql/controller"
)

func main() {
	r := gin.Default()
	r.GET("/posts", controller.ListPosts)
	r.POST("/posts", controller.CreatePost)
	r.GET("/posts/:id", controller.GetPostById)
	r.PUT("/posts/:id", controller.UpdatePostById)
	r.DELETE("/posts/:id", controller.DeletePostById)

	r.Run(":8089")
}
