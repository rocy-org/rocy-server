package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB

func init() {
	var err error
	username := os.Getenv("ROCY_SERVER_USERNAME")
	password := os.Getenv("ROCY_SERVER_PASSWORD")
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			//SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel: logger.Info, // Log level
		},
	)
	fmt.Println("username", username)
	fmt.Println("password", password)
	if username == "" || password == "" {
		panic("invalid username or password")
	}
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/rocy_server_db?charset=utf8mb4&parseTime=True&loc=Local", username, password)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

}
