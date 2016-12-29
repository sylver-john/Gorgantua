package utils

import (
	"log"
	"encoding/json"
	"errors"
	"io/ioutil"
	"../entity"	
)

func GetConfig(configPath string) (entity.Config, error) {
	var config entity.Config
	log.Print(config.Request.Base)
	rawConfig, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Print(err)
		return config, errors.New("Configuration file not found: " + configPath)
	}

	err = json.Unmarshal(rawConfig, &config)
	if err != nil {
		return config, errors.New("Unable to parse JSON configuration: " + err.Error())
	}
	return config, nil
}