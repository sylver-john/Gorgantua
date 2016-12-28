package mysql

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"math/rand"
	"regexp"
	"../entity"
	"../utils"
	"strings"
	"time"
)

var intValue = regexp.MustCompile(`int.*`)
var stringValue = regexp.MustCompile(`text`)
var varcharValue = regexp.MustCompile(`varchar`)
var dateValue = regexp.MustCompile(`date`)

func GetMysqlConnexion(config entity.Config) *sql.DB {
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

func GetMysqlColumns(db *sql.DB) []entity.MysqlParam {
	rows, err := db.Query("SHOW columns FROM table_test")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	params := make([]entity.MysqlParam, 0)
	for rows.Next() {
		var param entity.MysqlParam
		err := rows.Scan(&param.Field, &param.Coltype, &param.Null, &param.Key, &param.Coldefault, &param.Extra)
		if err != nil {
			log.Fatal(err)
		}
		params = append(params, param)
	}
	return params
}

func GenerateRow(params []entity.MysqlParam) map[string]entity.MysqlGeneratedData {
	mysqlGeneratedRowData := make(map[string]entity.MysqlGeneratedData, 0)
	var mysqlGeneratedData entity.MysqlGeneratedData
	for _, param := range params {
		switch {
		case stringValue.MatchString(param.Coltype):
			mysqlGeneratedData.Value = utils.RandSeq(10)
		case varcharValue.MatchString(param.Coltype):
			tmpArray := strings.Split(param.Coltype, "(")
			tmpString := strings.Replace(tmpArray[1], ")", "", -1)
			tmpInt, _:= strconv.Atoi(tmpString)
			mysqlGeneratedData.Value = utils.RandSeq(tmpInt)
		case intValue.MatchString(param.Coltype):
			mysqlGeneratedData.Value = rand.Intn(11)
		case dateValue.MatchString(param.Coltype):
			mysqlGeneratedData.Value = time.Now()
		}
		mysqlGeneratedRowData[param.Field] = mysqlGeneratedData
	}
	return mysqlGeneratedRowData
}

func GenerateQuery(request entity.Request, params []entity.MysqlParam, mysqlGeneratedData map[string]entity.MysqlGeneratedData) string {
	query := "INSERT " + request.Table + " SET "
	for _, param := range params {
		// il ne peut pas y avoir deux fois le même nom donc on peut utiliser cette méthode
		isLastElement := params[len(params)-1].Field == param.Field
		switch mysqlGeneratedData[param.Field].Value.(type) {
		case string:
			if !isLastElement {
				query += param.Field + "=\"" + mysqlGeneratedData[param.Field].Value.(string) +"\","
			} else {
				query += param.Field + "=\"" + mysqlGeneratedData[param.Field].Value.(string) + "\""			
			}
		case int:
			if !isLastElement {
				query += param.Field + "=" + strconv.Itoa(mysqlGeneratedData[param.Field].Value.(int)) +","
			} else {
				query += param.Field + "=" + strconv.Itoa(mysqlGeneratedData[param.Field].Value.(int))			
			}
		case time.Time:
			if !isLastElement {
				query += param.Field + "=" + mysqlGeneratedData[param.Field].Value.(time.Time).String() +","
			} else {
				query += param.Field + "=" + mysqlGeneratedData[param.Field].Value.(time.Time).String()
			}
		}
	}
	return query
}
func InsertMysql(request entity.Request, params []entity.MysqlParam, mysqlGeneratedData map[string]entity.MysqlGeneratedData, db *sql.DB) {
	query := GenerateQuery(request, params, mysqlGeneratedData)
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

func ExecuteAction(request entity.Request, params []entity.MysqlParam, db *sql.DB) {
	switch request.Action {
	case "INSERT":
		for i := 0; i < int(request.HowMany); i++ {
			mysqlGeneratedData := GenerateRow(params)
			InsertMysql(request, params, mysqlGeneratedData, db)
		}
	}
	defer db.Close()
}
