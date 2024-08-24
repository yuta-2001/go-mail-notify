resource "aws_ssm_parameter" "github_token" {
  name = "/${var.project_prefix}/github_token"
  type = "SecureString"
  value = "github_token"

  # valueはaws cliから上書きするため、変更を無視する
  lifecycle {
    ignore_changes = [value]
  }
}

resource "aws_ssm_parameter" "line_notify_token" {
  name = "/${var.project_prefix}/line_notify_token"
  type = "SecureString"
  value = "line_notify_token"

  # valueはaws cliから上書きするため、変更を無視する
  lifecycle {
    ignore_changes = [value]
  }
}
