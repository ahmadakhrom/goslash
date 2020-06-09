package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
)

var DB *gorm.DB

//connection for general connection
func init() {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/gormdb?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	DB = db
}
