package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Load config file
func Load() {
	// TODO: load different configs
	viper.SetConfigFile("config.yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("unable to read config: %v\n", err)
		os.Exit(1)
	}
}

// Get specific key in config
func Get(key string) interface{} {
	return viper.Get(key)
}
