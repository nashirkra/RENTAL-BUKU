package conf

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/nashirkra/RENTAL-BUKU/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// create new connection
func SetupDBConn() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load environment")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// dont forget to adjust your port
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3310)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to Database")
	}
	// model here

	db.AutoMigrate(&entity.User{})
	return db
}

// close the connection
func CloseDBConn(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from databse")
	}
	dbSQL.Close()
}
