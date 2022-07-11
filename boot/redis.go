package boot

import (
	"context"
	"fmt"
	"go-gin-api/pkg/config"
	"go-gin-api/pkg/console"
	"go-gin-api/pkg/g"
	"go-gin-api/pkg/redis"
)

type redisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

var Redis = redisConfig{}

func (r *redisConfig) Init() {
	config.UnmarshalKey("redis", &Redis)
	redis.NewClient(fmt.Sprintf("%v:%v", r.Host, r.Port), r.Password, r.DB)
	if err := g.Redis().Ping(context.Background()).Err(); err != nil {
		console.ExitIf(err)
	}
}
