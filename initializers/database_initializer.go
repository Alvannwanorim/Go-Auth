package initializers

import (
	"log"
	"os"

	"github.com/alvannwanorim/go-auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("error connecting the db")
	}
	log.Println("database connection successful...")

	SyncDatabase()
}

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
