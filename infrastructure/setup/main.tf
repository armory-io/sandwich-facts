provider "aws" {
  region = "us-east-1"
}

resource "aws_s3_bucket" "spinnaker-summit-state" {
  bucket_prefix = "armory-spinnaker-summit-"
  tags = {
    "armory.io/application" = "spinnaker-summit",
    "armory.io/owner"       = "engineering@armory.io"
    "armory.io/description" = "Bucket to store intermediate state for some Spinnaker Summit demos."
  }
}

output "tfstate_bucket" {
  value = aws_s3_bucket.spinnaker-summit-state.bucket
}
