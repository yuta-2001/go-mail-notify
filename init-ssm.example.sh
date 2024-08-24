#!/bin/bash

export AWS_DEFAULT_REGION=ap-northeast-1

aws ssm put-parameter --endpoint "http://localhost:4566" --name "/no-commit-notify/github_token" --type SecureString --value "**********"
aws ssm put-parameter --endpoint "http://localhost:4566" --name "/no-commit-notify/line_notify_token" --type SecureString --value "**********"
