package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go-curd/initializers"
	"go-curd/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func SignUp(c *gin.Context) {

	var err error

	//获取email和password从参数体
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "读取参数错误",
		})
		return
	}

	//获取密码hash值
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "获取密码hash错误",
		})
	}

	//创建用户
	user := models.User{
		Email:    body.Email,
		Password: string(hash),
	}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "创建用户失败",
		})
		return
	}
	//响应
	c.JSON(http.StatusOK, gin.H{
		"success": "创建用户成功",
	})
}

func Login(c *gin.Context) {

	//获取email和password从参数体
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "读取参数错误",
		})
		return
	}
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "密码或邮箱错误1",
		})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "密码或邮箱错误2",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "创建token失败",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
