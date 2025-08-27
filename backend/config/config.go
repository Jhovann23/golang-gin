package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	log.Println("Warning: no file .env found")
	if err != nil {
	}
}

func GetEnv(key string, defaultValue string) string {
	value, exist :=  os.LookupEnv(key)
	if !exist {
		return defaultValue
	}
	return value
}