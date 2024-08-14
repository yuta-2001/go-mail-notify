resource "aws_kms_key" "kms" {
  description = "Key for secret encryption and decryption"
  is_enabled  = true
}

resource "aws_kms_alias" "kms_alias" {
  name          = "alias/lambda-secret-key"
  target_key_id = aws_kms_key.kms.key_id
}
