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

    userName := os.Getenv("USER_GITHUB")
    if userName == "" {
        return "", "", "", fmt.Errorf("USER_GITHUB is empty")
    }

    githubToken := os.Getenv("TOKEN_GITHUB")
    if githubToken == "" {
        return "", "", "", fmt.Errorf("TOKEN_GITHUB is empty")
    }

    lineToken := os.Getenv("TOKEN_LINE_NOTIFY")
    if lineToken == "" {
        return "", "", "", fmt.Errorf("TOKEN_LINE_NOTIFY is empty")
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

