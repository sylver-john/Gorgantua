package main

import(
	"os"
	"log"
	"encoding/json"
	"errors"
	"io/ioutil"
	"./mysql"
	"./entity"
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


func main() {
	config, configErr := GetConfig("config.json")
	if configErr != nil {
		log.Fatal(configErr)
	}
	if len(os.Args) == 0 {
		log.Fatal("eu il manque un argument bonhomme")
	} else {
		switch os.Args[1] {
		case "mysql":
			db := mysql.GetMysqlConnexion(config)
			params := mysql.GetMysqlColumns(db)
			mysql.ExecuteAction(config.Request, params, db)
		}
	}
}

