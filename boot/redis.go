package boot

import (
	"context"
	"fmt"
	"go-gin-api/pkg/config"
	"go-gin-api/pkg/console"
	"go-gin-api/pkg/g"
	"go-gin-api/pkg/redis"
)

func initRedis() {
	var r redis.Config
	config.UnmarshalKey("redis", &r)
	redis.NewClient(fmt.Sprintf("%v:%v", r.Host, r.Port), r.Password, r.DB)
	if err := g.Redis().Ping(context.Background()).Err(); err != nil {
		console.ExitIf(err)
	}
}
