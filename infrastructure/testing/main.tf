variable "aws_region" {}
variable "account_id" {}

terraform {
  # Provided as part of Spinnaker pipeline.
  backend "s3" {}
}

provider "aws" {
  region = var.aws_region
  assume_role {
    role_name    = "arn:aws:iam::${var.account_id}:role/OrganizationAccountAccessRole"
    session_name = "EphemeralDemoResources"
  }
}

resource "aws_s3_bucket" "lunch_pail" {
  bucket_prefix = "sandwich-facts-"
  tags = {
    "armory.io/application" = "sandwich-facts",
    "armory.io/owner"       = "engineering@armory.io"
    "armory.io/description" = "Test bucket used to demonstrate ephemeral environment creation in Spinnaker pipelines."
  }
}

output "s3_bucket" {
  value = aws_s3_bucket.lunch_pail.bucket
}
