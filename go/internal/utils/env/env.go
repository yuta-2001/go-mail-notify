package env

import (
    "os"
    "fmt"
)

func GetEnv() (string, string, string, string, error) {
    env := os.Getenv("ENV")
    if env == "" {
        return "", "", "", "", fmt.Errorf("ENV is empty")
    }

    userName := os.Getenv("GITHUB_USER_NAME")
    if userName == "" {
        return "", "", "", "", fmt.Errorf("GITHUB_USER_NAME is empty")
    }

    githubToken := os.Getenv("GITHUB_TOKEN")
    if githubToken == "" {
        return "", "", "", "", fmt.Errorf("GITHUB_TOKEN is empty")
    }

    lineToken := os.Getenv("LINE_NOTIFY_TOKEN")
    if lineToken == "" {
        return "", "", "", "", fmt.Errorf("LINE_NOTIFY_TOKEN is empty")
    }

    return env, userName, githubToken, lineToken, nil
}
