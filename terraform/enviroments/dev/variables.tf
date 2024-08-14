variable "project_prefix" {
  default = "no-commit-notify"
}
variable "github_user" {
  sensitive = true
}
variable "github_token" {
  sensitive = true
}
variable "line_notify_token" {
  sensitive = true
}
