#!/bin/bash

export AWS_DEFAULT_REGION=ap-notheast-1

aws ssm put-parameter --endpoint "http://localhost:4566" --name "/no-commit-notify/token_github" --type SecureString --value "**********"
aws ssm put-parameter --endpoint "http://localhost:4566" --name "/no-commit-notify/token_line_notify" --type SecureString --value "**********"
