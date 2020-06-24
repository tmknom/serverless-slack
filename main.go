package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(event events.CloudWatchEvent) {
	log.Printf("event: %s", event)
	log.Printf("env.SLACK_WEB_HOOK: %s", os.Getenv("SLACK_WEB_HOOK"))
}

func main() {
	lambda.Start(handler)
}
