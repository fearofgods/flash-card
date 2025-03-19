package initialize

import (
	"fmt"

	"flash-card.juslt.click/global"
	"github.com/spf13/viper"
)

func LoadConfig(env string) {
	viper := viper.New()
	viper.SetConfigType("yaml")
	viper.SetConfigName(env)
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error on parsing env configuration file! %w", err))
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}
