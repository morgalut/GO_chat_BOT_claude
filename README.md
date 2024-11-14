 

# Claude Chatbot Project

This project is a command-line chatbot in Go that interacts with the Claude API from Anthropic. It includes a modular structure with separate files for API communication, logging, configuration, and data models.

## Prerequisites

To run this project, make sure you have the following:

- **Go** version 1.16 or later installed.
- **Claude API credentials** with the necessary permissions.

## Setting Up the Project

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/yourusername/claude-chatbot
   cd claude-chatbot
   ```

2. **Install Dependencies**:
   Ensure you are in the project root and run:
   ```bash
   go mod tidy
   ```

3. **Create an `.env` File**:
   To configure the chatbot, you need to create an `.env` file in the project’s root directory. This file should contain the API key, API URL, and model version. Here’s a template of what your `.env` file should look like:

   ```plaintext
   CLAUDE_API_KEY=sk-ant-api03-***********************************
   CLAUDE_API_URL=https://api.anthropic.com/v1/messages
   CLAUDE_MODEL=claude-3-5-sonnet-20241022
   ```

   Replace the asterisks (`*`) with your actual Claude API key.

4. **Run the Chatbot**:
   With your `.env` file set up, you can start the chatbot by running:

   ```bash
   go run cmd/chatbot/main.go
   ```

   The chatbot will start, and you can type messages to interact with it. Type `exit` to quit the application.

## Project Structure

- **cmd/chatbot/main.go**: The main entry point for the chatbot.
- **internal/api/client.go**: Handles API communication with the Claude API.
- **internal/config/config.go**: Loads configuration values from the `.env` file.
- **internal/logger/logger.go**: Sets up logging to record errors and debugging information.
- **pkg/models/chatbot.go**: Defines the data models for the chatbot.

## Important Notes

- **Environment Variables**: The `.env` file is critical to the chatbot’s functionality. Without this file, the program will fail to load the required API credentials and endpoint.
- **Error Handling**: The project includes structured logging for errors, making it easy to debug issues by checking the `logs/errors.log` file.


