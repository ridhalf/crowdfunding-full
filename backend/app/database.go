package app

import (
	"crowdfunding/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func NewDB() *gorm.DB {

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)
	return db
}
