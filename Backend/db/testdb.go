package db

import (
	"fmt"
	"os"

	"github.com/Samswrld02/Portfolio/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// create new db connection
func NewDb() *gorm.DB {
	fmt.Println("start db connection")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("DB_ADDRESS"), os.Getenv("MYSQL_DATABASE"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("cant connect %s", err.Error())
	}
	if err := db.AutoMigrate(&models.Project{}); err != nil {
		fmt.Printf("%s", err.Error())
	}
	if err := db.AutoMigrate(&models.User{}); err != nil {
		fmt.Printf("%s", err.Error())
	}

	return db
}
