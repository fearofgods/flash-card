package initialize

import (
	"flash-card.juslt.click/global"
	"go.uber.org/zap"
)

func RunApp() {
	env := LoadEnv()
	LoadConfig(env)
	appInfo := global.Config.AppInfo
	InitBanner(appInfo)
	InitLogger()
	global.Logger.Info("Log configuration is OK!", zap.String("ok", "success"))

	r := InitRouter()
	r.Run(global.Config.Server.Port)
}
