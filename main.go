package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(event events.CloudWatchEvent) {
	log.Printf("event: %s", event)
}

func main() {
	lambda.Start(handler)
}
