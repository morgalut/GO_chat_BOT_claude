package main

import (
    "fmt"
    "log"
    "strings"
    "my_chatbot_project/internal/api"
    "my_chatbot_project/internal/logger"
    "my_chatbot_project/internal/config"
)

const (
    modelVersion = "claude-3-5-sonnet-20241022" // Define model version here for easier updates
    maxTokens    = 1024                         // Set max tokens as a constant
)

func main() {
    // Load configuration (API key and URL)
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading configuration: %v", err)
    }

    // Initialize the logger
    logger, err := logger.CreateLogger("logs/errors.log")
    if err != nil {
        log.Fatalf("Could not set up logger: %v", err)
    }

    fmt.Println("Claude Chatbot is running. Type 'exit' to quit.")
    for {
        fmt.Print("You: ")
        var userInput string
        fmt.Scanln(&userInput)

        // Check for exit command or empty input
        if strings.ToLower(userInput) == "exit" {
            fmt.Println("Goodbye!")
            break
        }
        if strings.TrimSpace(userInput) == "" {
            fmt.Println("Please enter a message.")
            continue
        }

        // Send user input to the Claude API
        response, err := api.SendMessage(cfg.ApiURL, cfg.ApiKey, userInput, modelVersion, maxTokens, logger)
        if err != nil {
            fmt.Println("Sorry, something went wrong. Please try again.")
            continue
        }

        fmt.Println("Claude:", response)
    }
}
