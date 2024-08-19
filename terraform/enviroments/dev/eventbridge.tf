resource "aws_iam_role" "scheduler_role" {
  name = "EventBridgeSchedulerRole"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "scheduler.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy" "eventbridge_invoke_policy" {
  name = "EventBridgeInvokeLambdaPolicy"
  role = aws_iam_role.scheduler_role.id
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Action" : [
          "lambda:InvokeFunction"
        ],
        "Effect" : "Allow",
        "Resource" : "${aws_lambda_function.lambda.arn}"
      }
    ]
  })
}

resource "aws_scheduler_schedule" "invoke_lambda_schedule" {
  name = "InvokeLambdaSchedule"
  flexible_time_window {
    mode = "OFF"
  }
  schedule_expression = "cron(0 21 * * ? *)"
  target {
    arn      = aws_lambda_function.lambda.arn
    role_arn = aws_iam_role.scheduler_role.arn
    input    = "{}"
  }
}
