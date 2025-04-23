package testutil

import (
    "log"
    "github.com/joho/godotenv"
)

// LoadEnv loads the .env file for tests
func LoadEnv() {
    err := godotenv.Load("../.env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
}
