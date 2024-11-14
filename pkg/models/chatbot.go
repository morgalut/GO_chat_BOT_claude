 
package models

type ChatbotRequest struct {
    Input string `json:"input"`
}

type ChatbotResponse struct {
    Message string `json:"message"`
}
