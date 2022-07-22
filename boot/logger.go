package boot

import (
	"go-gin-api/pkg/config"
	log "go-gin-api/pkg/logger"
)

func initLogger() {
	var cfg log.Config
	config.UnmarshalKey("logger", &cfg)
	log.InitLogger(&cfg)
}
