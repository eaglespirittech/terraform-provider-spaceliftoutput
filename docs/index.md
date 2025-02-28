---
page_title: "Provider: Spacelift Output"
description: |-
  The Spacelift Output provider allows Terraform to retrieve outputs from Spacelift stacks.
---

# Spacelift Output Provider

The Spacelift Output provider allows Terraform to retrieve outputs from Spacelift stacks. This provider is useful when you need to use outputs from a Spacelift stack in your Terraform configuration.

## Example Usage

```terraform
terraform {
  required_providers {
    spaceliftoutput = {
      source = "eaglespirittech/spaceliftoutput"
      version = "0.1.0"
    }
  }
}

provider "spaceliftoutput" {
  api_token = "your-spacelift-api-token" # or use SPACELIFT_API_TOKEN env var
  account_name = "your-account-name" # optional, defaults to eaglespirittech or use spacelift_account_name env var
  # api_url = "https://your-account.app.spacelift.io/graphql" # optional
}

# Get all outputs from a stack
data "spaceliftoutput_stack_outputs" "example" {
  stack_id = "your-stack-id"
}

output "all_stack_outputs" {
  value = data.spaceliftoutput_stack_outputs.example.outputs
}

# Get a specific output from a stack
data "spaceliftoutput_stack_output" "example" {
  stack_id    = "your-stack-id"
  output_name = "output_name"
}

output "specific_output" {
  value = data.spaceliftoutput_stack_output.example.value
}
```

## Authentication

The Spacelift Output provider requires a Spacelift API token to authenticate with the Spacelift API. You can provide this token in one of two ways:

1. Set the `api_token` attribute in the provider configuration.
2. Set the `SPACELIFT_API_TOKEN` environment variable.

## Schema

### Optional

- **api_token** (String, Sensitive) - The Spacelift API token. Can also be set with the `SPACELIFT_API_TOKEN` environment variable.
- **account_name** (String) - Your account name in Spacelift. Used to construct the API URL if api_url is not specified. Defaults to `eaglespirittech`. Can also be set with the `spacelift_account_name` environment variable.
- **api_url** (String) - The Spacelift API URL. If not specified, it will be constructed using the account_name. 