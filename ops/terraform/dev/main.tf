# main.tf

provider "aws" {
  region                      = "us-east-1"  # Use any default region
  access_key                  = "test"  # Dummy credentials for Localstack
  secret_key                  = "test"
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true
  endpoints {
    sns = "http://localhost:4566"
    sqs = "http://localhost:4566"
  }
}

resource "aws_sns_topic" "user_events" {
  name = "user-events"
}