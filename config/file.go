package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

//LoadEnvFile Load config from the file that path is saved in os env.
func LoadFromEnvFile() {
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

//LoadCustomFromFile Load custom config from file, save to custom config
func LoadCustomFromFile(fileName string, customCfg interface{}) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := yaml.NewDecoder(f).Decode(customCfg); err != nil {
		return err
	}
	return nil
}
