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
  name               = local.lambda_iam_role_name
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
          "logs:*"
        ],
        Effect   = "Allow",
        Resource = "arn:aws:logs:*:${local.account_id}:log-group:/aws/lambda/${aws_lambda_function.lambda.function_name}:*"
      },
      {
        Action = [
          "ecr:GetDownloadUrlForLayer",
          "ecr:BatchGetImage",
          "ecr:BatchCheckLayerAvailability"
        ],
        Effect   = "Allow",
        Resource = "${aws_ecr_repository.repository.arn}"
      },
      { 
        Action = [
          "ssm:GetParameter"
        ],
        Effect   = "Allow",
        Resource = [
          "${aws_ssm_parameter.token_github.arn}",
          "${aws_ssm_parameter.token_line_notify.arn}"
        ]
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
  image_uri     = "${local.image_uri}:latest"
  timeout       = 30
  memory_size   = 128

  image_config {
    command = ["/main"]
  }

  architectures = ["arm64"]

  environment {
    variables = local.enviroment_variables
  }

  lifecycle {
    ignore_changes = [
      image_uri
    ]
  }
}
