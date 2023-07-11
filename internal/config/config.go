package config

import "github.com/spf13/viper"

type Config struct {
	Server `mapstructure:",squash"`
}

type Server struct {
	Port string `mapstructure:"SERVER_PORT"`
}

func New() (*Config, error) {

	viper.SetDefault("SERVER_PORT", "8080")

	viper.AutomaticEnv()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
