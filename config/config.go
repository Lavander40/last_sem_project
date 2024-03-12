package config

import "github.com/spf13/viper"

type Config struct {
	POSTGRES_USER     string `mapstructure:"POSTGRES_USER"`
	POSTGRES_PASSWORD string `mapstructure:"POSTGRES_PASSWORD"`
	POSTGRES_DB       string `mapstructure:"POSTGRES_DB"`

	NEO4J_USER     string `mapstructure:"NEO4J_USER"`
	NEO4J_PASSWORD string `mapstructure:"NEO4J_PASSWORD"`
}

func InitConfig() (*Config, error) {
	var config Config

	viper.SetConfigFile(".env")
	viper.SetConfigName("")
	viper.AddConfigPath(".")
	// viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
