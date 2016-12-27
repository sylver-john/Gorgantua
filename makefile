setup:
	@cp src/conf.json.dist src/conf.json
	@go get github.com/go-sql-driver/mysql
	@echo "Setup completed"
run:
	go run src/main.go mysql