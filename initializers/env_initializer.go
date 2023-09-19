package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Panic("error loading .env")
	}
}
