package config

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/apollo.v0"
	"gopkg.in/yaml.v3"
)

//LoadEnvFile Load config from the file that path is saved in os env.
func LoadFromEnvFile() {
	loadConfigsFromYamlFile()
	if IsApolloEnable() {
		if err := apollo.StartWithConf(&config.Apollo.Conf); err != nil {
			fmt.Printf("load apollo config error: %s\n", err)
			os.Exit(1)
			return
		}

		go func() {
			for {
				event := apollo.WatchUpdate()
				changeEvent := <-event
				bytes, _ := json.Marshal(changeEvent)
				fmt.Println("event:", string(bytes))
			}
		}()
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

func loadConfigsFromYamlFile() {
	appCfgFileName := os.Getenv(appConfigFileEnvKey)
	if err := LoadFromFile(appCfgFileName); err != nil {
		fmt.Printf("load file config error: %s\n", err)
		os.Exit(1)
	}

	apiCfgFileName := os.Getenv(apiConfigFileEnvKey)
	if err := LoadFromFile(apiCfgFileName); err != nil {
		fmt.Printf("load file config error: %s\n", err)
		os.Exit(1)
	}

}
