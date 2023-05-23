terraform {
  backend "s3" {
    bucket = "tfstate-prod2"
    key    = "terraform-state/terraform.tfstate"
    region = "eu-central-1"
  }
}

# Dynamo DB to lock terraform



