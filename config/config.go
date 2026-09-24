package config

import (
	"github.com/spf13/viper"
)

type Conf struct {
	WeatherAPIKey string `mapstructure:"weather_api_key"`
	HTTPPort      string `mapstructure:"http_port"`
}

func LoadConfig(path string) (*Conf, error) {
	var cfg Conf

	viper.SetConfigFile(path + "/.env")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
