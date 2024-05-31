package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	DB *gorm.DB
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	var errDb error
	DB, errDb = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errDb != nil {
		panic("failed to connect database")
	}

	// Auto Migrate
	// DB.AutoMigrate(&Model{}) // Uncomment and replace Model with your model struct if needed

	fmt.Println("Database connected successfully")
}

// GetDB returns a pointer to the database instance
func GetDB() *gorm.DB {
	return DB
}
