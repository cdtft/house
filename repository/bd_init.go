package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

const jdbcUrl = "root:root@tcp(127.0.0.1:3306)/house?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

func init() {
	db, err := gorm.Open(mysql.Open(jdbcUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Panicln(err)
	}
	DB = db
}
