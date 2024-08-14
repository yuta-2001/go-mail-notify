package main

import (
    "fmt"

    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/kms"

    "no-commit-notify/go/internal/github"
    "no-commit-notify/go/internal/line"
    "no-commit-notify/go/internal/env"
)

// HandleRequest is the entry point for AWS Lambda function
func HandleRequest() {
    
    sess, err := session.NewSession()
    if err != nil {
        fmt.Println(err)
        return
    }
    svc := kms.New(sess)

    userName, githubToken, lineToken, err := env.GetEnv(svc)
    if err != nil {
        fmt.Println(err)
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
