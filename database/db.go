package database

import (
	"OctavianoRyan25/GO-JWT/model"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// host     = "localhost"
	// port     = "5432"
	// user     = "postgres"
	// password = "mulianet010"
	// dbname   = "api"
	
	db       *gorm.DB
)

func StartDB() {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Gagal memuat file .env")
    }
	
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")
    dbName := os.Getenv("DB_NAME")
	config := fmt.Sprintf("host = %s user = %v password = %v dbname = %v port = %v sslmode = disable", dbHost, dbUser, dbPass, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	db.Debug().AutoMigrate(model.User{}, model.Photo{}, model.SocialMedia{}, model.Comment{})
}

func GetDB() *gorm.DB {
	return db
}