package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
)

// LoadConfig load the configuration file and return Config object.
func LoadConfig() (config Config, err error) {
	var configName = "config.json"

	if len(os.Args) > 1 {
		configName = os.Args[1]
	}

	file, err := os.Open(configName)
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

// RouterConfig load the configuration file and return RouterURI object for router endpoint's
func RouterConfig() (config RouterURI, err error) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return RouterURI{}, err
	}

	var configName = path.Join(pwd, "routers", "router.json")

	if len(os.Args) > 2 {
		configName = os.Args[2]
	}

	file, err := os.Open(configName)
	if err != nil {
		return RouterURI{}, err
	}
	defer file.Close()

	value, err := io.ReadAll(file)
	if err != nil {
		return RouterURI{}, err
	}

	err = json.Unmarshal(value, &config)
	if err != nil {
		return RouterURI{}, err
	}

	return config, nil
}
