package main

import(
	"os"
	"log"
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Config struct {
	Host string `json:"host"`
	User string `json:"user"`
	Password string `json:"password"`
	Request	Request `json:"request"`
}

type Request struct {
	Base string `json:"base"`
	Table string `json:"table"`
	Action string `json:"ation"`
	Params []Param `json:"params"`
	HowMany float64 `json:"howMany"` 
}

type Param struct {
	Name string `json:"name"`
	Type string `json:type`
}

func GetConfig(configPath string) (Config, error) {
	var config Config
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
			GetMysqlConnexion(config)
		}
	}
}

