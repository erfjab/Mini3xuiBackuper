package config

import (
	"errors"
	"github.com/spf13/viper"
)

type Config struct {
	PanelUsername	 string `mapstructure:"PANEL_USERNAME"`
	PanelPassword	 string `mapstructure:"PANEL_PASSWORD"`
	PanelHost 	     string `mapstructure:"PANEL_HOST"`
}

var Cfg *Config

func LoadConfig() (*Config, error) {
	var config Config

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		var configNotFound viper.ConfigFileNotFoundError
		if !errors.As(err, &configNotFound) {
			return &config, err
		}
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		return &config, err
	}

	if err = config.validate(); err != nil {
		return &config, err
	}

	Cfg = &config
	return &config, nil
}

func (c *Config) validate() error {
	if c.PanelUsername == "" {
		return errors.New("PANEL_USERNAME is required")
	}
	if c.PanelPassword == "" {
		return errors.New("PANEL_PASSWORD is required")
	}
	if c.PanelHost == "" {
		return errors.New("PANEL_HOST is required")
	}
	return nil
}

