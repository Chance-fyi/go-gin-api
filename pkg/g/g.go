package g

import (
	"go-gin-api/pkg/database"
	"gorm.io/gorm"
)

func DB(name ...string) *gorm.DB {
	return database.DB.Connect(name...)
}
