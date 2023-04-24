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
