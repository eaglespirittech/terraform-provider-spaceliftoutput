terraform {
  required_providers {
    spaceliftoutput = {
      source = "eaglespirittech/spaceliftoutput"
    }
  }
}

provider "spaceliftoutput" {
  # Configuration options
  # api_token = "your-spacelift-api-token" # or use SPACELIFT_API_TOKEN env var
  # account_name = "your-account-name" # optional, defaults to eaglespirittech or use spacelift_account_name env var
  # api_url = "https://your-account.app.spacelift.io/graphql" # optional, constructed from account_name if not provided
} 