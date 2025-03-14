package config

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"github.com/joho/godotenv"
)

// Dbinstance holds the Gorm database connection
type Dbinstance struct {
	Db *gorm.DB
}

// DB is a global variable to hold the database connection
var DB Dbinstance

// ConnectDb initializes the database connection using Gorm
func ConnectDb() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbType := os.Getenv("DB_TYPE")
	if dbType == "" {
		log.Fatal("DB_TYPE environment variable is not set")
	}

	// Use the ConnectionURLBuilder utility to get the DSN
	dsn, err := ConnectionURLBuilder(dbType)
	if err != nil {
		log.Fatalf("Failed to build connection URL: %v", err)
	}

	// Initialize Gorm connection
	var db *gorm.DB
	switch dbType {
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: gormLogger.Default.LogMode(gormLogger.Info),
		})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
			Logger: gormLogger.Default.LogMode(gormLogger.Info),
		})
	
	default:
		log.Fatalf("Unsupported database type: %s", dbType)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	//db.AutoMigrate(&model.UserModel{})

	log.Println("Connected to database successfully using Gorm!")

	// Assign the database connection to the global DB variable
	DB = Dbinstance{
		Db: db,
	}
}
