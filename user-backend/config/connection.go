package config

import (
	"go-api/models/entity"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
	err error
)

func Connect() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
		  SlowThreshold:              time.Second,   // Slow SQL threshold
		  LogLevel:                   logger.Info, // Log level
		  IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
		  Colorful:                  false,          // Disable color
		},
	  )
	// docker  
	//  dsn := "root:db@2012@tcp(host.docker.internal:3306)/user-backend?charset=utf8mb4&parseTime=True&loc=Local"

	// local
	dsn := "root:db@2012@tcp(localhost:3306)/user-backend?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&entity.User{}, &entity.Role{}, &entity.Product{})
	return DB
}