package env

import (
    "encoding/base64"
    "os"
    "fmt"

    "github.com/aws/aws-sdk-go/service/kms"
)

func GetEnv(svc *kms.KMS) (string, string, string, error) {
    env := os.Getenv("ENV")
    if env == "" {
        return "", "", "", fmt.Errorf("ENV is empty")
    }

    userName := os.Getenv("GITHUB_USER")
    if userName == "" {
        return "", "", "", fmt.Errorf("GITHUB_USER is empty")
    }

    githubToken := os.Getenv("GITHUB_TOKEN")
    if githubToken == "" {
        return "", "", "", fmt.Errorf("GITHUB_TOKEN is empty")
    }

    lineToken := os.Getenv("LINE_LIFF_TOKEN")
    if lineToken == "" {
        return "", "", "", fmt.Errorf("LINE_LIFF_TOKEN is empty")
    }

    if env != "local" {
        var err error
        userName, err = decryptVariable(svc, userName)
        if err != nil {
            return "", "", "", err
        }

        githubToken, err = decryptVariable(svc, githubToken)
        if err != nil {
            return "", "", "", err
        }

        lineToken, err = decryptVariable(svc, lineToken)
        if err != nil {
            return "", "", "", err
        }
    }

    return userName, githubToken, lineToken, nil
}


func decryptVariable(svc *kms.KMS, variable string) (string, error) {
    if variable == "" {
        return "", fmt.Errorf("variable is empty")
    }

    v, err := base64.StdEncoding.DecodeString(variable)
    if err != nil {
        return "", err
    }

    input := kms.DecryptInput{
        CiphertextBlob: v,
    }

    result, err := svc.Decrypt(&input)
    if err != nil {
        return "", err
    }

    return string(result.Plaintext), nil
}

