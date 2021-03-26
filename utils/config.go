package utils

import (
	"log"

	"github.com/spf13/viper"
)

// ReadConfig is a method to read yaml
func ReadConfig(cfgPath string, cfgName string) (cfg *viper.Viper) {
	cfg = viper.New()
	cfg.AddConfigPath(cfgPath)
	cfg.SetConfigName(cfgName)
	cfg.SetConfigType("yaml")

	if err := cfg.ReadInConfig(); err != nil {
		log.Fatal("utils.config.go->ReadConfig error:", err.Error())
		return nil
	}
	return cfg
}
