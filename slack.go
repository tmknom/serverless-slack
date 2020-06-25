package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var (
	IncomingWebhookURL string = os.Getenv("SLACK_INCOMING_WEBHOOK_URL")
)

type SlackMessage struct {
	Text string `json:"text"`
}

func (sm *SlackMessage) post() error {
	raw, err := json.Marshal(sm)
	if err != nil {
		return fmt.Errorf("marshal failed: %s, input: %s", err, sm)
	}

	req, err := http.NewRequest(http.MethodPost, IncomingWebhookURL, bytes.NewReader(raw))
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
