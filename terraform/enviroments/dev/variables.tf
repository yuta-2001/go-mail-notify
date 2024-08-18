variable "project_prefix" {
  default = "no-commit-notify"
}
variable "enviroment" {
  default = "production"
}
variable "user_github" {
  sensitive = true
}
variable "token_github" {
  sensitive = true
}
variable "token_line_notify" {
  sensitive = true
}
