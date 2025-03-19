package initialize

import (
	"fmt"

	"flash-card.juslt.click/pkg/setting"
)

func InitBanner(info setting.AppInfo) {
	fmt.Println("========================================")
	fmt.Println("App name: ", info.Name)
	fmt.Println("Version: ", info.Version)
	fmt.Println("Env: ", info.Env)
	fmt.Println("========================================")
}
