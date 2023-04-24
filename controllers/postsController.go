package controllers

import (
	"github.com/gin-gonic/gin"
	"go-curd/initializers"
	"go-curd/models"
	"log"
)

func PostsCreate(c *gin.Context) {
	//创建body结构体
	var body struct {
		Body  string
		Title string
	}
	//Bind是用于将请求的参数绑定到一个结构体上的方法。具体来说，它可以将请求的参数解析并映射到一个结构体的字段上，然后将这个结构体作为请求处理函数的参数传入。
	err := c.Bind(&body)
	if err != nil {
		log.Fatal("bind body failed")
	}
	//创建一个post请求
	post := models.Post{Title: body.Title, Body: body.Body}
	//执行操作到数据库
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}
	//返回创建的数据
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//定义posts变量，对应查询结果初始列表
	var posts []models.Post
	//gorm查询列表语法。使用GORM库中的DB对象执行了查询操作，将查询结果赋值到了切片类型的posts变量中
	//其中，&posts表示获取posts变量的指针，这样查询结果可以直接被写入到切片变量中
	//在GORM中，Find方法可以用于查询数据库中的数据并返回查询结果，如果查询结果为空，则返回一个空的切片。
	//该方法接收一个指向切片的指针，以便能够直接将查询结果写入切片变量中。
	initializers.DB.Find(&posts)
	//返回查询结果
	c.JSON(200, gin.H{
		"post": posts,
	})
}

func PostsShow(c *gin.Context) {
	//从url中获取id
	id := c.Param("id")

	//创建post结构体
	var post models.Post
	//DB.First(&post, id)是用于检索与给定id相符的第一个记录，并将结果存储在post变量中。
	//它接受两个参数：第一个参数是指向变量的指针，第二个参数是要检索的记录的id
	//如果找到记录，则将其读入post中，并返回nil作为错误；
	//如果未找到记录，则将post设置为默认值（通常为0或nil），并返回ErrRecordNotFound作为错误。
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"posts": post,
	})
}
