package env

import (
    "os"
    "fmt"

    "no-commit-notify/go/internal/aws/ssm"
)

func GetEnv() (string, string, string, error) {
    env, err := fetchEnvVar("env")
    if err != nil {
        return "", "", "", err
    }

    userName, err := fetchEnvVar("github_user")
    if err != nil {
        return "", "", "", err
    }

    githubTokenParam, err := fetchEnvVar("github_token_param_name")
    if err != nil {
        return "", "", "", err
    }

    lineTokenParam, err := fetchEnvVar("line_notify_token_param_name")
    if err != nil {
        return "", "", "", err
    }

    ssmInstance, err := ssm.GetSsmInstance(env)
    if err != nil {
        return "", "", "", err
    }

    githubToken , err := ssmInstance.GetParamValue(githubTokenParam, true)
    if err != nil {
        return "", "", "", err
    }

    lineToken , err := ssmInstance.GetParamValue(lineTokenParam, true)
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
