package main

import (
    "no-commit-notify/go/githubhelper"
    "no-commit-notify/go/linehelper"
    "github.com/aws/aws-lambda-go/lambda"
    "fmt"
)

// HandleRequest is the entry point for AWS Lambda function
func HandleRequest() {
    contributesCount, err := githubhelper.GetContributesCount()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Contributes count:", contributesCount)

    err = linehelper.SendMessage(contributesCount)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Message sent successfully")
}

func main() {
    lambda.Start(HandleRequest)
}
