package env

import (
    "fmt"
    "os"

    "no-commit-notify/go/internal/aws/ssm"
)

func GetEnv() (string, string, string, error) {
    env, err := fetchEnvVar("ENV")
    if err != nil {
        return "", "", "", err
    }

    userName, err := fetchEnvVar("GITHUB_USER")
    if err != nil {
        return "", "", "", err
    }

    githubTokenParam, err := fetchEnvVar("GITHUB_TOKEN_PARAM_NAME")
    if err != nil {
        return "", "", "", err
    }

    lineTokenParam, err := fetchEnvVar("LINE_NOTIFY_TOKEN_PARAM_NAME")
    if err != nil {
        return "", "", "", err
    }

    ssmInstance, err := ssm.GetSsmInstance(env)
    if err != nil {
        return "", "", "", err
    }

    githubToken, err := ssmInstance.GetParamValue(githubTokenParam, true)
    if err != nil {
        return "", "", "", err
    }

    lineToken, err := ssmInstance.GetParamValue(lineTokenParam, true)
    if err != nil {
        return "", "", "", err
    }

    return userName, githubToken, lineToken, nil
}

func fetchEnvVar(varName string) (string, error) {
    value := os.Getenv(varName)
    if value == "" {
        return "", fmt.Errorf("%s is empty", varName)
    }
    return value, nil
}
