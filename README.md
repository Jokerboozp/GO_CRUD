## Go-CURD

### 一、初始化项目

#### 1.1 go设置代理

```go
$ go env -w GO111MODULE = on
$ go env -w GOPROXY = https: //goproxy.cn,direct
```

#### 1.2 安装CompileDaemon

```go
网址：https: //github.com/githubnemo/CompileDaemon
获取命令：go get github.com/githubnemo/CompileDaemon
安装命令：go install github.com/githubnemo/CompileDaemon
运行命令：CompileDaemon -command = "./go-curd" （./go -curd是项目目录）
```

#### 1.3 安装Gin

```go
网址：https: //gin-gonic.com/
命令：go get -u github.com/gin-gonic/gin
```

#### 1.4 安装GORM

```go
网址：https: //gorm.io/zh_CN/
安装：go get -u gorm.io/gorm
安装postgres数据库模块：go get -u gorm.io/driver/postgres
```

#### 1.5 安装godotenv

```go
网址：https: //github.com/joho/godotenv
获取：go get github.com/joho/godotenv
安装：go install github.com/joho/godotenv/cmd/godotenv@latest
```

#### 1.6 运行项目

- 复制Gin官网例子至main方法中，运行CompileDaemon，之后访问localhost:8080/ping。如果返回信息`{"message":"pong"}`，则项目启动成功

![iSJ3W8.png](https://i.328888.xyz/2023/04/24/iSJ3W8.png)
![iSJNP5.png](https://i.328888.xyz/2023/04/24/iSJNP5.png)
![iSJguC.png](https://i.328888.xyz/2023/04/24/iSJguC.png)

#### 1.7 设置.env

- 新建.env文件，项目设置端口号

```txt
PORT=3000
```

- main方法中引入.env（方法参照godoenv的GitHub参考文件）。

```go
func init() {
err := godotenv.Load()
if err != nil {
log.Fatal("Error loading .env file")
}
}
```

- 如果引入成功，项目会自动重新运行在3000端口
  ![iSJoud.png](https://i.328888.xyz/2023/04/24/iSJoud.png)
  ![iSJvHb.png](https://i.328888.xyz/2023/04/24/iSJvHb.png)

#### 1.8 设置.env目录

- 创建initializers文件夹，并在其中创建loadEnvVariables.go文件。

- 文件中设置LoadEnvVariables方法，并将main.go中的init函数搬过来

```go
package initializers

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
```

- 在main.go中的init方法中引入LoadEnvVariables方法

```go
func init() {
initializers.LoadEnvVariables()
}
```

- 如果没有错误，程序仍然会正确运行在3000端口，并可以通过`localhost:3000/ping`访问到

#### 1.9 连接数据库

- .env文件中添加数据库连接(连接格式为GORM确定的)

```txt
DB_URL = "host=localhost user=postgres password=root dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
```

- 创建database.go

```go
package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

// DB 声明一个名为 DB 的变量，其类型是 *gorm.DB，即指向 gorm.DB 类型的指针
// 通常情况下，这样的声明被用来在程序的全局范围内创建一个可以被其他函数和方法访问的全局变量。
// 在这个特定的例子中，DB 是一个与数据库连接相关的变量。
// 由于指针默认为 nil，因此在声明时没有初始化 DB，需要在程序的某个地方对它进行初始化，才能使用它
var DB *gorm.DB

func ConnectToDB() {
	//定义一个变量来存储异常
	var err error
	//GORM定义的连接数据库方法.使用godotenv获取DB_URL
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//判断错误
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}
```

- main.go中引用ConnectToDB方法

```go
func init() {
initializers.LoadEnvVariables()
initializers.ConnectToDB()
}
```

- 项目正常运行在3000端口并可正常访问`localhost:3000/ping`即为配置成功

#### 1.10 声明模型

- 创建models文件夹，并创建postModel.go。用来对应数据库中的字段

```go
package models

import "gorm.io/gorm"

// Post gorm定义了一个gorm.Model结构体
/*type Model struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
*/
//所以下面的struct是等同于
/**
type User struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
  Name string
}
*/
type Post struct {
	gorm.Model
	Title string
	Body  string
}
```

#### 1.11 migrate自动生成表

- AutoMigrate是一个用于自动创建数据库表的方法。它会根据结构体的定义创建数据库表，如果表已经存在则会检查每个字段是否存在，如果有缺失的字段则会自动添加。同时，如果在结构体中定义了新的字段，它也会自动添加到表中

```go
package main

import (
	"go-curd/initializers"
	"go-curd/models"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	//Gorm语法
	err := initializers.DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal("failed migrate database")
	}
}
```

- 运行migrate.go，自动生成Post模型对应的数据库表.

```go
//(在项目根目录运行，否则LoadEnvVariables方法会找不到.env文件)
go run migrate/migrate.go
```

- 运行成功会在对应数据库中发现表已创建

![iSPgDJ.png](https://i.328888.xyz/2023/04/24/iSPgDJ.png)

### 二、编写增加接口

#### 2.1 增加PostsCreate方法。

- postsController.go

```go
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
```

#### 2.2 在main.go中调用添加接口

- main.go

```go
func main() {
r := gin.Default()
//调用posts接口进行信息新增
r.POST("/posts", controllers.PostsCreate)
r.Run()
}
```

#### 2.3 接口测试

- 接口测试结果

![iSap4J.png](https://i.328888.xyz/2023/04/24/iSap4J.png)

### 三、编写查询所有信息列表接口

#### 3.1 增加PostsIndex方法

```go
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
```

#### 3.2 在main.go中调用查询方法

```go
r.GET("/posts", controllers.PostsIndex)
```

#### 3.3 接口测试

- 接口测试结果

![iS1rbk.png](https://i.328888.xyz/2023/04/24/iS1rbk.png)

### 四、编写按照ID查询信息接口

#### 4.1 增加PostsShow方法

```go
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
```

#### 4.2 在main.go中调用根据ID查询列表方法

```go
r.GET("/posts/:id", controllers.PostsShow)
```

#### 4.3 接口测试

- 测试结果

![iSQQ2H.png](https://i.328888.xyz/2023/04/24/iSQQ2H.png)