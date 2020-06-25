package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var incomingWebhookURL = os.Getenv("SLACK_INCOMING_WEBHOOK_URL")

func post(message interface{}) error {
	raw, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("marshal failed: %s, input: %#v", err, message)
	}

	req, err := http.NewRequest(http.MethodPost, incomingWebhookURL, bytes.NewReader(raw))
	if err != nil {
		return fmt.Errorf("failed new request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to post webhook: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %d", resp.StatusCode)
	}
	return nil
}
