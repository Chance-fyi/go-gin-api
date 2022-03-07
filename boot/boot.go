package boot

import "go-gin-api/pkg/config"

type boot struct{}

var Boot = boot{}

func (*boot) Init() {
	//初始化配置
	config.Init()
}
