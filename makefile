setup:
	@go get github.com/go-sql-driver/mysql
	@echo "Setup completed"
run:
	go run src/main.go src/mysql.go mysql