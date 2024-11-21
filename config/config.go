package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	User   string
	DbName string
	SSL    string
}

func InitDb() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &Config{
		User:   os.Getenv("USER"),
		DbName: os.Getenv("NAME"),
		SSL:    os.Getenv("SSL"),
	}
}
