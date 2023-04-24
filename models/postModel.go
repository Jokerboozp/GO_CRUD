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
