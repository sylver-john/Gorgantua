package entity

import(
	"database/sql"
)
type Config struct {
	Database string `json:"database"`
	Host string `json:"host"`
	User string `json:"user"`
	Password string `json:"password"`
	Request	Request `json:"request"`
}

type Request struct {
	Base string `json:"base"`
	Table string `json:"table"`
	Action string `json:"action"`
	HowMany float64 `json:"howMany"` 
}

type MysqlParam struct {
	Field	string
	Coltype	string
	Null	sql.NullString
	Key	sql.NullString
	Coldefault	sql.NullString
	Extra	sql.NullString
}

type MysqlGeneratedData struct {
	Value interface{}
}