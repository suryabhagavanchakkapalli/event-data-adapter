# Event Webhook Server

## Overview
This Go program serves as an HTTP server to receive JSON events, converts them into a structured format, and sends the data to a predefined webhook. It acts as a bridge between incoming events and an external service.

## Usage
1. Update the `WebhookURL` constant in the code with your desired webhook endpoint.
2. Run the program: `go run main.go`.
3. The server will be accessible at `http://localhost:8080/receive`.
4. Send JSON requests to the `/receive` endpoint.

## Data Structures
- **Attribute**: Represents an attribute with a value and type.
- **UserTrait**: Represents a user trait with a value and type.
- **EventData**: Represents the structured format for event data.

## Main Functions
- **main**: Sets up an HTTP server and starts a worker goroutine to process incoming requests concurrently.
- **worker**: Processes messages from a channel, converts JSON requests into `EventData`, and sends it to the webhook.
- **processRequest**: Converts a dynamic map from JSON requests into an `EventData` structure.
- **sendToWebhook**: Sends the converted `EventData` to a predefined webhook using an HTTP POST request.

## Configuration
- **WebhookURL**: The URL of the webhook. Modify this constant to use a different endpoint.

## Dependencies
- Standard Go libraries for HTTP handling and JSON processing.
