package main

import(
	"log"
	"./mysql"
	"./utils"
)

func main() {
	config, configErr := utils.GetConfig("config.json")
	if configErr != nil {
		log.Fatal(configErr)
	}
	switch config.Database {
	case "mysql":
		db := mysql.GetMysqlConnexion(config)
		params := mysql.GetMysqlColumns(db)
		mysql.ExecuteAction(config.Request, params, db)
	}
}

