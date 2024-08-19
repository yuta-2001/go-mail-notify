package env

import (
    "encoding/base64"
    "os"
    "fmt"
	"strings"

    "github.com/aws/aws-sdk-go/service/kms"
)

func GetEnv(svc *kms.KMS) (string, string, string, error) {
    env := os.Getenv("env")
    if env == "" {
        return "", "", "", fmt.Errorf("env is empty")
    }

    userName := os.Getenv("user_github")
    if userName == "" {
        return "", "", "", fmt.Errorf("user_github is empty")
    }

    githubToken := os.Getenv("token_github")
    if githubToken == "" {
        return "", "", "", fmt.Errorf("token_github is empty")
    }

    lineToken := os.Getenv("token_line_notify")
    if lineToken == "" {
        return "", "", "", fmt.Errorf("token_line_notify is empty")
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
		fmt.Println("Error in decryptVariable")
        return "", err
    }

    input := kms.DecryptInput{
        CiphertextBlob: v,
    }

    result, err := svc.Decrypt(&input)
    if err != nil {
        return "", err
    }

    return strings.Replace(string(result.Plaintext), "\n", "", -1), nil
}
