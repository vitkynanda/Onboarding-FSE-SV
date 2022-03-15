package seeder

import (
	"go-api/models/entity"
	"testing"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	errDB error
)

func TestSeedRoles(t *testing.T) {

	dsn := "root:@tcp(127.0.0.1:3306)/user-backend?charset=utf8mb4&parseTime=True&loc=Local"

	DB, errDB = gorm.Open(mysql.Open(dsn))

	if errDB != nil {
		t.Fatal("failed to connect database")
	}

	roles := []entity.Role{
		{ID: uuid.New().String(), Title: "admin", Active: true},
		{ID: uuid.New().String(), Title: "maker", Active: true},
		{ID: uuid.New().String(), Title: "checker", Active: true},
		{ID: uuid.New().String(), Title: "signer", Active: true},
		{ID: uuid.New().String(), Title: "viewer", Active: true},
	}

	err := DB.Create(&roles).Error
	if err != nil {
		t.Fatal("error batch insert to mysql", err.Error())
	}
}
