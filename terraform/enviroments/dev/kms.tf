resource "aws_kms_key" "kms" {
  description = "Key for secret encryption and decryption"
  is_enabled  = true
}

resource "aws_kms_alias" "kms_alias" {
  name          = "alias/lambda-secret-key"
  target_key_id = aws_kms_key.kms.key_id
}

resource "aws_kms_ciphertext" "github_user" {
  key_id    = aws_kms_key.kms.key_id
  plaintext = local.secrets.github_user
}

resource "aws_kms_ciphertext" "github_token" {
  key_id    = aws_kms_key.kms.key_id
  plaintext = local.secrets.github_token
}

resource "aws_kms_ciphertext" "line_notify_token" {
  key_id    = aws_kms_key.kms.key_id
  plaintext = local.secrets.line_notify_token
}
