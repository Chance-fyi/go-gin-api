package boot

import "go-gin-api/pkg/config"

func Init() {
	//初始化配置
	config.Init()
	initDatabase()
	initRedis()
	initLogger()
}
