package config

import (
	"github.com/spf13/viper"
)

const configFile = "gometr"

type Config struct {
	HTTPHost string
	HTTPPort string
}

func NewConfig() (*Config, error) {
	viper.SetConfigName(configFile) // lookup gometr.yaml
	viper.AddConfigPath("configs")
	viper.AddConfigPath("Service/configs")

	if err := viper.ReadInConfig(); err != nil {
		return &Config{}, err
	}

	c := &Config{
		HTTPHost: viper.GetString("http.host"),
		HTTPPort: viper.GetString("http.port"),
	}

	return c, nil
}
