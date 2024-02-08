package server

import (
	"encoding/json"
	"event-data-adapter/data"
	"event-data-adapter/utils"
	"fmt"
	"io"
	"net/http"
)

// Worker function to process messages from the channel
func Worker(channel chan map[string]interface{}) {
	for {
		// Receive message from the channel
		dynamicMap := <-channel

		// Process the request and convert it into the desired format
		eventData := data.ProcessRequest(dynamicMap)

		// Print the converted event (you can modify this part to do something else with the converted event)
		eventDataJSON, _ := json.MarshalIndent(eventData, "", "  ")
		fmt.Println(string(eventDataJSON))

		// Send the converted event to the webhook
		utils.SendToWebhook(eventData)
	}
}

// HandleRequest handles HTTP requests
func HandleRequest(w http.ResponseWriter, r *http.Request, channel chan map[string]interface{}) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var dynamicMap map[string]interface{}
	err = json.Unmarshal(body, &dynamicMap)
	if err != nil {
		http.Error(w, "Error unmarshaling JSON", http.StatusBadRequest)
		return
	}

	// Send the request to the worker via the channel
	channel <- dynamicMap

	w.WriteHeader(http.StatusOK)
}
