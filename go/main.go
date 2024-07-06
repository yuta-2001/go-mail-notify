package main

import (
	"no-commit-notify/go/githubhelper"
	"no-commit-notify/go/linehelper"
	"github.com/aws/aws-lambda-go/lambda"
	"fmt"
)

// HandleRequest is the entry point for AWS Lambda function
func HandleRequest() {
	contributesCount := githubhelper.GetContributesCount()
	linehelper.SendMessage(contributesCount)
	fmt.Println("Contributes count: ", contributesCount)
	fmt.Println("Message sent to Line.")
}

func main() {
	lambda.Start(HandleRequest)
}
