package database

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var db *gorm.DB

func Connect() {
	_, currentFile, _, _ := runtime.Caller(0)
	projectDir := filepath.Dir(currentFile)
	envFile := filepath.Join(projectDir, "../../.env")
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("PORT")
	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + port + " sslmode=disable TimeZone=Asia/Tashkent"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = database
}

/*func Connect() {
	dsn := "host=localhost user=postgres password=zaxscdvfbgnhmjqwerty dbname=myshop port=5432 sslmode=disable TimeZone=Asia/Tashkent"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil { panic(err) }
	db = database
}*/

func GetDB() *gorm.DB { return db }
