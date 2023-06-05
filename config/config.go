package config

import (
	"encoding/json"
	"io"
	"os"
)

// LoadConfig load the configuration file with viper package
func LoadConfig() (config Config, err error) {
	file, err := os.Open("config.json")
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	value, err := io.ReadAll(file)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(value, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
