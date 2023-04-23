package config

import (
	"github.com/spf13/viper"
)

// LoadConfig load the configuration file with viper package
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config") // name of config file
	viper.SetConfigType("json")   // file type extension

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)

	return
}
