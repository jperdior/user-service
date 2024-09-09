package main

import (
	"golang-template/cmd/api/bootstrap"
	_ "golang-template/docs"
	"log"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
