package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Port string `json:Port`

	DBHostname string `json:DBHostname`
	DBPort     string `json:DBPort`
	Username   string `json:Username`
	Password   string `json:Password`
}

// New creates a new, global Config. Reads in configuration from config files.
func New() (*Config, error) {
	cfg := &Config{}

	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	v.SetConfigName("conf")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		fmt.Println("No config file found; only reading from environment")
	}

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
