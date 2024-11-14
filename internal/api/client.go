package api

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "bytes"
    "my_chatbot_project/pkg/models"
)

func SendMessage(apiURL, apiKey, input, modelVersion string, maxTokens int, logger *log.Logger) (string, error) {
    requestBody := map[string]interface{}{
        "model": modelVersion,
        "max_tokens": maxTokens,
        "messages": []map[string]string{
            {
                "role":    "user",
                "content": input,
            },
        },
    }

    reqBody, err := json.Marshal(requestBody)
    if err != nil {
        logger.Printf("Failed to create request body: %v\n", err)
        return "", fmt.Errorf("failed to create request body: %v", err)
    }

    logger.Printf("Request Body: %s\n", string(reqBody))

    req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(reqBody))
    if err != nil {
        logger.Printf("Failed to create HTTP request: %v\n", err)
        return "", fmt.Errorf("failed to create HTTP request: %v", err)
    }

    req.Header.Set("x-api-key", apiKey)
    req.Header.Set("anthropic-version", "2023-06-01")
    req.Header.Set("Content-Type", "application/json")
    logger.Printf("Headers: x-api-key=%s, anthropic-version=2023-06-01, Content-Type=application/json\n", apiKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        logger.Printf("Error sending request: %v\n", err)
        return "", fmt.Errorf("error sending request: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        body, _ := io.ReadAll(resp.Body)
        var apiError map[string]interface{}
        if err := json.Unmarshal(body, &apiError); err == nil {
            if msg, ok := apiError["error"].(map[string]interface{})["message"].(string); ok {
                logger.Printf("API error message: %s\n", msg)
                return "", fmt.Errorf("API error: %s", msg)
            }
        }
        logger.Printf("Non-200 response: %d\nResponse Body: %s\n", resp.StatusCode, string(body))
        return "", fmt.Errorf("non-200 response: %d", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        logger.Printf("Failed to read response: %v\n", err)
        return "", fmt.Errorf("failed to read response: %v", err)
    }

    var chatbotResponse models.ChatbotResponse
    if err := json.Unmarshal(body, &chatbotResponse); err != nil {
        logger.Printf("Failed to parse response: %v\n", err)
        return "", fmt.Errorf("failed to parse response: %v", err)
    }

    return chatbotResponse.Message, nil
}
