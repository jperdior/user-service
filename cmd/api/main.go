package main

import (
	"log"
	"user-service/cmd/api/bootstrap"
	_ "user-service/docs"
	"user-service/internal/platform/database/mysql"
)

func main() {

	mysql.ConnectDB()
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
