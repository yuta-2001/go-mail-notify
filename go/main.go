package main

import (
    "fmt"

    "github.com/aws/aws-lambda-go/lambda"

    "no-commit-notify/go/internal/github"
    "no-commit-notify/go/internal/line"
	"no-commit-notify/go/internal/utils/env"
)

// HandleRequest is the entry point for AWS Lambda function
func HandleRequest() {
	env, userName, githubToken, lineToken, err := env.GetEnv()
	if err != nil {
		fmt.Println(err)
		return
	}

	if env == "local" {
		fmt.Println("Running locally")
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
