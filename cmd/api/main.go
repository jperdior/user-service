package main

import (
	"log"
	"user-service/cmd/api/bootstrap"
	_ "user-service/docs"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
