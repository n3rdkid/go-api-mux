package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//GetEnvByKey -> Get value of key in env
func GetEnvByKey(key string) string {
	return os.Getenv(key)
}

//LoadEnv -> Load the env file initially
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}
