package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

//loadFromEnv load config file from env
func loadFromEnv() {
	fileName := os.Getenv(configFileEnvKey)
	if err := LoadFromFile(fileName); err != nil {
		fmt.Printf("load file config error: %s\n", err)
		os.Exit(1)
	}
}

//LoadFromFile from file
func LoadFromFile(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := yaml.NewDecoder(f).Decode(&config); err != nil {
		return err
	}
	return nil
}
