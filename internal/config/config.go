package config

import (
    "os"
    "fmt"
    "github.com/joho/godotenv"
)

type Config struct {
    ApiURL string
    ApiKey string
}

func LoadConfig() (*Config, error) {
    // Load variables from .env file
    if err := godotenv.Load(); err != nil {
        return nil, fmt.Errorf("error loading .env file: %v", err)
    }

    apiURL := os.Getenv("CLAUDE_API_URL")
    apiKey := os.Getenv("CLAUDE_API_KEY")

    if apiURL == "" || apiKey == "" {
        return nil, fmt.Errorf("missing required environment variables")
    }

    return &Config{
        ApiURL: apiURL,
        ApiKey: apiKey,
    }, nil
}
