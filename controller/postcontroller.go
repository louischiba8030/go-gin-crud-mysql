package controller

import (
	"net/http"
//	"context"
	"fmt"
	"log"
	"strconv"
	"github.com/gin-gonic/gin"

	"go-gin-crud-mysql/model"
	"go-gin-crud-mysql/util"
)

func GetPostById(c *gin.Context){
	p := model.Post{}
	p.FirstById(c.Param("id"))

	c.JSONP(http.StatusOK, p)
}

func ListPosts(c *gin.Context){
	//name := c.PostForm("name")
	p := []model.Post{}
	model.DB.Find(&p)

	c.JSONP(http.StatusOK, p)
}

func CreatePost (c *gin.Context) {
	ap, _ := strconv.Atoi(c.PostForm("age"))
	age := uint(ap)

	post := model.Post {
		Ulid: util.GenerateULID(),
		Name: c.PostForm("name"),
		Age: age,
		Bloodtype: c.PostForm("bloodtype"),
		Origin: c.PostForm("origin"),
	}

	fmt.Print("added post: ", post)
	post.Create()

	c.JSONP(http.StatusCreated, post)
}

func UpdatePostById (c *gin.Context) {
	ap, _ := strconv.Atoi(c.PostForm("age"))
	age := uint(ap)

	post := model.Post {
		Name: c.PostForm("name"),
		Age: age,
		Bloodtype: c.PostForm("bloodtype"),
		Origin: c.PostForm("origin"),
	}
	post.Updates(c.Param("id"))
	fmt.Println("updated post: ", post)

	c.JSON(http.StatusOK, post)
}

func DeletePostById (c *gin.Context) {
	p := model.Post{}
	p.DeleteById(c.Param("id"))

	log.Println("Deleted post: ", p)

	c.JSON(http.StatusOK, p)
}
