package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jakskal/user-login/cmd/router"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	pathDir := "cd ../../"
	err := godotenv.Load(filepath.Join(pathDir, ".env"))
	fmt.Println(err)
	if err != nil {
		log.Fatal("Error loading .env file")
		log.Fatal(err)
	}
}

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	db, err := gorm.Open("postgres", `host=`+dbHost+` port=`+dbPort+` user=`+dbUser+` dbname=`+dbName+` password=`+dbPassword+` sslmode=disable`)
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}
	defer db.Close()
	db.LogMode(true)
	handler := initHandler(db)
	router.API(*handler)
}
