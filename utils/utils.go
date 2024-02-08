package utils

import (
	"bytes"
	"encoding/json"
	"event-data-adapter/data"
	"fmt"
	"net/http"
	"sync"
)

var mutex sync.Mutex

const WebhookURL = "https://webhook.site/212d98c8-cc06-4f15-9f2a-f8e9e1350f06"

// SendToWebhook sends converted event to the webhook
func SendToWebhook(convertedEvent data.EventData) {
	fmt.Println("webhookURL:", WebhookURL)
	mutex.Lock()
	defer mutex.Unlock()

	webhookData, err := json.Marshal(convertedEvent)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println("Webhook data:", string(webhookData)) // Print the webhook data

	// Make an HTTP POST request to the webhook
	response, err := http.Post(WebhookURL, "application/json", bytes.NewBuffer(webhookData))
	if err != nil {
		fmt.Println("Error sending data to webhook:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Webhook returned non-OK status:", response.Status)
		// Handle the error or log it based on your requirements
		return
	}

	fmt.Println("Webhook response:", response.Status)
}
