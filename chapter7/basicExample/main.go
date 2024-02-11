package main

import (
	"go_trial/gorest/chapter7/basicExample/helper"
	"log"
)

func main() {
	_, err := helper.InitDB()
	if err != nil {
		log.Println(err)
	}
	log.Println("Database tables are successfully initialized.")
}
