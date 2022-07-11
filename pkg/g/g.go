package g

import (
	"github.com/go-redis/redis/v9"
	"go-gin-api/pkg/database"
	r "go-gin-api/pkg/redis"
	"gorm.io/gorm"
)

func DB(name ...string) *gorm.DB {
	return database.DB.Connect(name...)
}

func Redis() *redis.Client {
	return r.Redis
}
