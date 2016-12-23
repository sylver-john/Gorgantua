package main

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetMysqlConnexion(config Config) {
		db, err := sql.Open("mysql", config.User + ":" + config.Password + "@tcp(" + config.Host +")")
		if err != nil {
			log.Fatal(err)
		}
		err = db.Ping()
		if err != nil {
			log.Print("yolo")
		}
		defer db.Close()
}