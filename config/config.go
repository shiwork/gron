package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type GronConfig struct {

}

func Parse(filename string) (GronConfig, error) {
	var config GronConfig
	jsonString, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Failed to read config file:", err)
		return config, err
	}

	err = json.Unmarshal(jsonString, &config)
	if err != nil {
		log.Println("Failed to json unmarshal:", err)
		return config, nil
	}
	return config, nil
}
