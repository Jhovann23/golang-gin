package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: no file .env found")
	}
}

func GetEnv(key string, defaultValue string) string {
	value, exist :=  os.LookupEnv(key)
	if !exist {
		return defaultValue
	}
	return value
}