package initialize

import (
	"flash-card.juslt.click/global"
	"flash-card.juslt.click/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
