package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(event events.CloudWatchEvent) error {
	fmt.Printf("event: %s", event)
	message := map[string]interface{}{"text": "posted by lambda"}
	return post(message)
}

func main() {
	lambda.Start(handler)
}
