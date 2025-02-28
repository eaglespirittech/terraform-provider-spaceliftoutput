# Terraform Provider for Spacelift Stack Outputs

This Terraform provider allows you to retrieve outputs from Spacelift stacks and use them in your Terraform configurations.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.21

## Building The Provider

1. Clone the repository
2. Enter the repository directory
3. Build the provider using the Go `install` command:

```shell
go install
```

## Using the provider

```hcl
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

## Example Usage

```hcl
# Get all outputs from a Spacelift stack
data "spaceliftoutput_stack_outputs" "my_stack" {
  stack_id = "my-stack-id"
}

# Use a specific output in another resource
resource "aws_instance" "example" {
  ami           = data.spaceliftoutput_stack_outputs.my_stack.outputs["ami_id"]
  instance_type = "t2.micro"
}

# Get a specific output directly
data "spaceliftoutput_stack_output" "vpc_id" {
  stack_id    = "my-stack-id"
  output_name = "vpc_id"
}

# Use the specific output in another resource
resource "aws_security_group" "example" {
  vpc_id = data.spaceliftoutput_stack_output.vpc_id.value
  # ...
}
```

## Development

### Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.21

### Building

1. Clone the repository
2. Enter the repository directory
3. Build the provider using the Go `install` command:

```shell
go install
```

### Testing

To run the tests, execute:

```shell
go test -v ./...
```

## Publishing to Terraform Registry

This provider can be published to the Terraform Registry by following these steps:

1. Create a GitHub release with a semantic version tag (e.g., v0.1.0)
2. The GitHub Actions workflow will automatically build and publish the provider to the Terraform Registry

