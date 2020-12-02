package utils

import (
	"os"

	"github.com/joho/godotenv"
)

//ReadEnv function is for read .env file
func ReadEnv(key string, defaultVal string) string {
	godotenv.Load(".env")
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultVal
	}
	return value
}
