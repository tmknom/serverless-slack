package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(event events.CloudWatchEvent) error {
	fmt.Printf("event: %s", event)
	fmt.Printf("env.SLACK_WEB_HOOK: %s", os.Getenv("SLACK_WEB_HOOK"))
	return nil
}

func main() {
	lambda.Start(handler)
}
