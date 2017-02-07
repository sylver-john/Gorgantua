# Gorgantua
Fill a database with some random insert writing in Go

## Setup
1. Install Go: obviously you should install Go if you want to run the script
2. Fill src/config.dist/json with your parameters like :
```json
{
	"database": "mysql",
	"host": "127.0.0.1:3306",
	"user": "root",
	"password": "",
	"request": {
		"base": "base_test",
		"table": "table_test",
		"action": "INSERT",
		"howMany": 50
	}
}
```
3. If you have make you can  use `make setup and then make run`, if not you have to install a package first with `get github.com/go-sql-driver/mysql`, fill src/config.json instead of src/config.dist.json and finally run `go run src/main.go`

## Configuration

* `database`: only mysql supported currently
* `host`: your server's address with the port
* `user` and `password`: to connect to your database
* `request.base`: the base where your targeted table is
* `request.table`: the table you want to fill
* `request.action`: only insert supported currently
* `request.howMany`: the number of rows which will be inserted
