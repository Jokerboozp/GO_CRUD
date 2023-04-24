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
