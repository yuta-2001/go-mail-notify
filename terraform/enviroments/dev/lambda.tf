data "aws_iam_policy_document" "lambda_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "lambda_execution_role" {
  name = local.lambda_iam_role_name
  assume_role_policy = data.aws_iam_policy_document.lambda_assume_role_policy.json
}

resource "aws_iam_policy" "lambda_policy" {
  name        = local.lambda_iam_policy_name
  description = "IAM policy for Lambda function"
  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ],
        Effect   = "Allow",
        resources = ["arn:aws:logs:*:*:*"]
      },
      {
        Action = [
          "ecr:GetDownloadUrlForLayer",
          "ecr:BatchGetImage",
          "ecr:BatchCheckLayerAvailability"
        ],
        Effect   = "Allow",
        Resource = "*"
      },
      {
        Action = [
          "kms:Decrypt"
        ]
        Effect   = "Allow"
        Resource = "*"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_policy_attachment" {
  policy_arn = aws_iam_policy.lambda_policy.arn
  role       = aws_iam_role.lambda_execution_role.name
}

resource "aws_lambda_function" "lambda" {
  function_name = local.function_name
  role          = aws_iam_role.lambda_execution_role.arn
  package_type  = "Image"
  image_uri     = "${local.image_uri}:firstImage"
  timeout       = 30
  memory_size   = 128

  image_config {
    command = ["/main"]
  }

  environment {
    variables = {
      GITHUB_USER       = "yuta-2001"
      GITHUB_TOKEN      = ""
      LINE_NOTIFY_TOKEN = ""
    }
  }

  architectures = ["arm64"]

  # kms_key_arn = var.kms_key_arn

  lifecycle {
    ignore_changes = [
      image_uri
    ]
  }
}
