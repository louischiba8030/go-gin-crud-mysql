package model

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-gin-crud-mysql/config"
)

var DB *gorm.DB
var err error

func init() {
	// Connect to DB
	conf := config.Config
	dsn := conf.DbUser + ":" + conf.DbPassword + "@tcp(" + conf.DbHost + ":" + conf.DbPort + ")/" + conf.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	// Open DB
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(dsn + "database can't connect")
	}

	// 
	// DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	// Make migration
	DB.AutoMigrate(&Post{})
}
