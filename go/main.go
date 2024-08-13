package main

import (
    "os"
    "fmt"

    "github.com/aws/aws-lambda-go/lambda"

    "no-commit-notify/go/internal/github"
    "no-commit-notify/go/internal/line"
)

// HandleRequest is the entry point for AWS Lambda function
func HandleRequest() {
    userName := os.Getenv("GITHUB_USER")
    githubToken := os.Getenv("GITHUB_TOKEN")
    lineToken := os.Getenv("LINE_NOTIFY_TOKEN")

    if userName == "" || githubToken == "" || lineToken == "" {
        fmt.Println("please set variables")
        return
    }

    contributesCount, err := github.GetContributesCount(userName, githubToken)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Contributes count:", contributesCount)

    err = line.SendMessage(contributesCount, lineToken)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Message sent successfully")
}

func main() {
    lambda.Start(HandleRequest)
}
