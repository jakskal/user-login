package main

import (
	"log"

	"github.com/jakskal/user-login/cmd/router"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=example password=mypassword sslmode=disable")
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}
	defer db.Close()
	db.LogMode(true)
	handler := initHandler(db)
	router.API(*handler)
}
