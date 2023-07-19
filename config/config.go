package config

import (
	"github.com/spf13/viper"
)

type Lala struct {
	DBUser string
	DBPass string
	DBHost string
	DBName string
	DBCharset string
}


func GetEnvConfig() {
	viper.AddConfigPath("./")
	viper.SetConfigType("env")
	viper.SetConfigName("env")
	viper.ReadInConfig()
}