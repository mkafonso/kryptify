package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DB_DRIVER              string        `mapstructure:"DB_DRIVER"`
	DB_SOURCE              string        `mapstructure:"DB_SOURCE"`
	GRPC_SERVER_ADDRESS    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TOKEN_SYMETRIC_KEY     string        `mapstructure:"TOKEN_SYMETRIC_KEY"`
	ACCESS_TOKEN_DURATION  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	REFRESH_TOKEN_DURATION time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func LoadEnvironmentVariable(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
