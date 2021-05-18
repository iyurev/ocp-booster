package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	appName           = "ocp-booster"
	configFileName    = "config.toml"
	defaultConfigPath = "/etc/ocp-booster"
)

func NewConfig() (*Cluster, error) {
	var cluster Cluster
	config := viper.New()
	//Config file
	config.SetConfigName(configFileName)
	config.AddConfigPath(defaultConfigPath)
	config.AddConfigPath(".")
	config.AddConfigPath(fmt.Sprintf("$HOME/%s", appName))
	config.SetConfigType("toml")
	if err := config.ReadInConfig(); err != nil {
		return nil, err
	}
	err := config.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &cluster, nil
}
