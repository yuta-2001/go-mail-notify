package env

import (
    "os"
    "fmt"

    "no-commit-notify/go/internal/kms"
)

func GetEnv() (string, string, string, error) {
    env := os.Getenv("env")
    if env == "" {
        return "", "", "", fmt.Errorf("env is empty")
    }

    userName := os.Getenv("user_github")
    if userName == "" {
        return "", "", "", fmt.Errorf("user_github is empty")
    }

    kmsInstance, err := kms.GetKmsInstance(env)
    if err != nil {
        return "", "", "", err
    }

    githubToken := os.Getenv("token_github")
    if githubToken == "" {
        return "", "", "", fmt.Errorf("token_github is empty")
    }
    githubToken, err = kmsInstance.DecryptVariable(githubToken)
    if err != nil {
        return "", "", "", err
    }

    lineToken := os.Getenv("token_line_notify")
    if lineToken == "" {
        return "", "", "", fmt.Errorf("token_line_notify is empty")
    }
    lineToken, err = kmsInstance.DecryptVariable(lineToken)
    if err != nil {
        return "", "", "", err
    }

    return userName, githubToken, lineToken, nil
}
