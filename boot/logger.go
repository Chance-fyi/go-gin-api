package boot

import (
	"go-gin-api/pkg/config"
	log "go-gin-api/pkg/logger"
)

type logger struct {
}

var Logger = logger{}

func (*logger) Init() {
	var cfg log.Config
	config.UnmarshalKey("logger", &cfg)
	log.InitLogger(&cfg)
}
