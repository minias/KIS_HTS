// internal/config/loader.go
package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// LoadConfig load config by profile
func LoadConfig() (*Config, error) {

	env := os.Getenv("APP_ENV")

	if env == "" {
		env = "dev"
	}

	viper.SetConfigName("config." + env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	viper.AutomaticEnv()

	var cfg Config

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	fmt.Println("Loaded profile:", env)

	return &cfg, nil
}
