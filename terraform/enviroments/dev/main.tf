terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  backend "s3" {
    bucket  = "no-commit-notify-terraform-state"
    key     = "state/dev/terraform.tfstate"
    region  = "ap-northeast-1"
    encrypt = true
  }
}

data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

locals {
  account_id                  = data.aws_caller_identity.current.account_id
  region                      = data.aws_region.current.name
  bucket_name                 = "${var.project_prefix}-${local.account_id}"
  function_name               = var.project_prefix
  lambda_iam_role_name        = "${var.project_prefix}-iam-lambda-role"
  lambda_iam_policy_name      = "${var.project_prefix}-iam-lambda-policy"
  eventbridge_iam_role_name   = "${var.project_prefix}-iam-eventbridge-role"
  eventbridge_iam_policy_name = "${var.project_prefix}-iam-eventbridge-policy"
  repository_name             = var.project_prefix
  image_uri                   = "${local.account_id}.dkr.ecr.${local.region}.amazonaws.com/${local.repository_name}"
  schedule_rule               = "${var.project_prefix}-schedule-rule"
}

provider "aws" {
  region = local.region
}

module "iam" {
  source                      = "../../modules/iam"
  lambda_iam_policy_name      = local.lambda_iam_policy_name
  lambda_iam_role_name        = local.lambda_iam_role_name
}

module "kms" {
  source = "../../modules/kms"
}

module "ecr" {
  source          = "../../modules/ecr"
  repository_name = local.repository_name
}

module "cloudwatch" {
  source = "../../modules/cloudwatch_log_group"
  function_name = local.function_name
}

module "lambda" {
  source        = "../../modules/lambda"
  function_name = local.function_name
  image_uri     = local.image_uri
  iam_role_arn  = module.iam.lambda_role_arn
  kms_key_arn   = module.kms.kms_key_arn
}

module "eventbridge" {
  source        = "../../modules/eventbridge"
  schedule_rule = local.schedule_rule
  lambda_arn    = module.lambda.lambda_arn
  function_name = local.function_name
}
