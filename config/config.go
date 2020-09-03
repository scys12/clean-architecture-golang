package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Secret            string `mapstructure:"secret"`
	DBDriver          string `mapstructure:"db_driver"`
	DBName            string `mapstructure:"db_name"`
	DBHost            string `mapstructure:"db_host"`
	DBPort            string `mapstructure:"db_port"`
	JwtAccessExpires  int    `mapstructure:"jwt_at_expire"`
	JwtRefreshExpires int    `mapstructure:"jwt_rt_expire"`
}

func parseConfigFilePath() string {
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return workDir
}

func NewConfig() *Config {
	configPath := parseConfigFilePath()
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s\n", err))
	}
	config := new(Config)
	if err := viper.Unmarshal(config); err != nil {
		panic(fmt.Errorf("failed to parse config file: %w\n", err))
	}
	return config
}
