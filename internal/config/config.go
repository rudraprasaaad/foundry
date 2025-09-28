package config

import "github.com/spf13/viper"

type Config struct {
	ServerPort string `mapstructure:"port"`
	JWTSecret  string `mapstructure:"secret"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
