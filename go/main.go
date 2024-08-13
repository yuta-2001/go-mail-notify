package main

import (
    "no-commit-notify/go/internal/github"
	"no-commit-notify/go/internal/line"
    "github.com/aws/aws-lambda-go/lambda"
    "fmt"
)

// HandleRequest is the entry point for AWS Lambda function
func HandleRequest() {
    contributesCount, err := github.GetContributesCount()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Contributes count:", contributesCount)

    err = line.SendMessage(contributesCount)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Message sent successfully")
}

func main() {
    lambda.Start(HandleRequest)
}
