package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func init() {
	pathDir := "cd ../../"
	err := godotenv.Load(filepath.Join(pathDir, ".env"))
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
	// dbPassword := os.Getenv("DB_PASSWORD")

	dbURL := `postgres://postgres:` + dbUser + `@` + dbHost + `:` + dbPort + `/` + dbName + `?sslmode=disable`
	m, err := migrate.New(
		"file://cmd/migrate/file",
		// "postgres://postgres:postgres@localhost:5432/example?sslmode=disable",
		dbURL,
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
