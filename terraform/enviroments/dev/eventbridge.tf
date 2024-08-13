
##### setting up IAM roles and policies for EventBridge
data "aws_iam_policy_document" "eventbridge_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["events.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "eventbridge_execution_role" {
  name = local.eventbridge_iam_role_name
  assume_role_policy = data.aws_iam_policy_document.eventbridge_assume_role_policy.json
}

resource "aws_cloudwatch_event_rule" "event_rule" {
  name        = "event_rule"
  description = "Event rule to trigger Lambda function every 5 minutes"
  schedule_expression = "cron(0 12 * * ? *)"
}

resource "aws_cloudwatch_event_target" "lambda_target" {
  rule      = aws_cloudwatch_event_rule.event_rule.name
  target_id = "my_lambda_target"
  arn       = aws_lambda_function.lambda.arn
  role_arn = aws_iam_role.eventbridge_execution_role
}

resource "aws_lambda_permission" "allow_eventbridge" {
  statement_id  = "AllowExecutionFromEventBridge"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda.name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.event_rule.arn
}
