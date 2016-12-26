package main

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	field	string
	coltype	string
	null	sql.NullString
	key	sql.NullString
	coldefault	sql.NullString
	extra	sql.NullString
)

type MysqlParam struct {
	field	string
	coltype	string
	null	sql.NullString
	key	sql.NullString
	coldefault	sql.NullString
	extra	sql.NullString
}

type MysqlGeneratedData struct {
	name string
	value interface{}
}

func GetMysqlConnexion(config Config) *sql.DB {
	db, err := sql.Open("mysql", config.User + ":" + config.Password + "@tcp(" + config.Host + ")/" + config.Request.Base)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetMysqlColumns(db *sql.DB) []MysqlParam {
	rows, err := db.Query("SHOW columns FROM table_test")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	params := make([]MysqlParam, 0)
	for rows.Next() {
		var param MysqlParam
		err := rows.Scan(&param.field, &param.coltype, &param.null, &param.key, &param.coldefault, &param.extra)
		if err != nil {
			log.Fatal(err)
		}
		params = append(params, param)
	}
	defer db.Close()
	return params
}

func GenerateData(params []MysqlParam) []MysqlGeneratedData {
	mysqlGeneratedData := make([]MysqlGeneratedData, 0)
	for _, param := range params {
		log.Print(param)
		switch param.coltype {
		case "text":
			// generate random string
		case "int":
			// generate random int
		}
	}
	return mysqlGeneratedData
}

func InsertMysql(mysqlGeneratedData []MysqlGeneratedData, db *sql.DB) {

}

func ExecuteAction(request Request, params []MysqlParam, db *sql.DB) {
	switch request.Action {
	case "INSERT":
		for i := 0; i < int(request.HowMany); i++ {
			mysqlGeneratedData := GenerateData(params)
			log.Print(mysqlGeneratedData)
			InsertMysql(mysqlGeneratedData, db)
			// on les insert
		}
	}
}
