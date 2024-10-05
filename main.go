package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func init() {
	// membuat environment variable ENV untuk menentukan environment
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	// memilih file .env berdasarkan environment
	envFile := fmt.Sprintf(".env.%s", env)

	// memuat file .env berdasarkan environment
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s: %v", envFile, err)
	}

}
func main() {
	// mendapatkan value environment variable
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// format connection string menggunakan variable dari .env
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// membuka koneksi menggunakan pgx
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	defer db.Close()

	// cek koneksi db
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	fmt.Println("Successfully connected to the database!")
}
