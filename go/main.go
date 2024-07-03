package main

import (
	"no-commit-notify/go/githubhelper"
	"no-commit-notify/go/linehelper"
	"github.com/aws/aws-lambda-go/lambda"
)

// HandleRequest is the entry point for AWS Lambda function
func HandleRequest() {
	contributesCount := githubhelper.GetContributesCount()
	linehelper.SendMessage(contributesCount)
}

func main() {
	lambda.Start(HandleRequest)
}
