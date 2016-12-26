package main

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"math/rand"
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
	return params
}

func GenerateRow(params []MysqlParam) map[string]MysqlGeneratedData {
	mysqlGeneratedRowData := make(map[string]MysqlGeneratedData, 0)
	var mysqlGeneratedData MysqlGeneratedData
	for _, param := range params {
		log.Print(param)
		switch param.coltype {
		case "text":
			// generate random string
			mysqlGeneratedData.value = "test"
		case "int(11)":
			// generate random int
			mysqlGeneratedData.value = rand.Intn(11)
		}
		mysqlGeneratedRowData[param.field] = mysqlGeneratedData
	}
	return mysqlGeneratedRowData
}

func InsertMysql(request Request, params []MysqlParam, mysqlGeneratedData map[string]MysqlGeneratedData, db *sql.DB) {
	query := "INSERT " + request.Table + " SET "
	for _, param := range params {
		// il ne peut pas y avoir deux fois le même nom donc on peut utiliser cette méthode
		isLastElement := params[len(params)-1].field == param.field
		switch mysqlGeneratedData[param.field].value.(type) {
		case string:
			if !isLastElement {
				query += param.field + "=\"" + mysqlGeneratedData[param.field].value.(string) +"\","
			} else {
				query += param.field + "=\"" + mysqlGeneratedData[param.field].value.(string) + "\""			
			}
		case int:
			if !isLastElement {
				query += param.field + "=" + strconv.Itoa(mysqlGeneratedData[param.field].value.(int)) +","
			} else {
				query += param.field + "=" + strconv.Itoa(mysqlGeneratedData[param.field].value.(int))			
			}
		}
	}
	log.Print(query)
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)	
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)	
	}
}

func ExecuteAction(request Request, params []MysqlParam, db *sql.DB) {
	switch request.Action {
	case "INSERT":
		for i := 0; i < int(request.HowMany); i++ {
			mysqlGeneratedData := GenerateRow(params)
			log.Print(mysqlGeneratedData)
			InsertMysql(request, params, mysqlGeneratedData, db)
			defer db.Close()
		}
	}
}
