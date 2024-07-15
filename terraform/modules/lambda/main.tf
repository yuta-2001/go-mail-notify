resource "aws_lambda_function" "lambda" {
  function_name = var.function_name
  role          = var.iam_role_arn
  package_type  = "Image"
  image_uri     = "${var.image_uri}:latest"
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

  kms_key_arn = var.kms_key_arn
}

